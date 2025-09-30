package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"melody_cure/DAO"
	"melody_cure/config"
	"melody_cure/model"
	"net/http"
	"strings"
	"time"
)

type AIReportService struct {
	generatedReportDAO *DAO.GeneratedReportDAO
	healingLogDAO      *DAO.HealingLogDAO
}

func NewAIReportService(generatedReportDAO *DAO.GeneratedReportDAO, healingLogDAO *DAO.HealingLogDAO) *AIReportService {
	return &AIReportService{
		generatedReportDAO: generatedReportDAO,
		healingLogDAO:      healingLogDAO,
	}
}

// GenerateReport 生成AI报告
func (s *AIReportService) GenerateReport(childArchiveID string, reportType string, startDate, endDate *time.Time) (*model.GeneratedReport, error) {
	// 将字符串ID转换为uint（这里假设childArchiveID是数字字符串）
	// 在实际项目中，可能需要根据具体的ID类型进行调整
	// 暂时使用模拟数据，实际项目中需要实现ID转换逻辑
	
	// 模拟获取疗愈记录（实际项目中需要根据ID类型调整）
	var logs []model.HealingLog
	// logs, err := s.healingLogDAO.GetHealingLogsByChildIDWithDateFilter(childID, startDate, endDate)
	// if err != nil {
	//     return nil, fmt.Errorf("获取疗愈记录失败: %v", err)
	// }

	// 构建AI请求内容
	content, err := s.buildAIPrompt(logs, reportType)
	if err != nil {
		return nil, fmt.Errorf("构建AI请求失败: %v", err)
	}

	// 调用AI API生成内容
	generatedContent, err := s.callAIAPI(content)
	if err != nil {
		return nil, fmt.Errorf("AI生成失败: %v", err)
	}

	// 创建报告记录
	report := &model.GeneratedReport{
		ChildArchiveID: childArchiveID,
		ReportType:     reportType,
		Content:        generatedContent,
		IsEdited:       false,
		GeneratedAt:    time.Now(),
	}

	err = s.generatedReportDAO.CreateGeneratedReport(report)
	if err != nil {
		return nil, fmt.Errorf("保存报告失败: %v", err)
	}

	return report, nil
}

// UpdateReportContent 更新报告内容
func (s *AIReportService) UpdateReportContent(reportID uint, content string) error {
	report := &model.GeneratedReport{
		ID:       reportID,
		Content:  content,
		IsEdited: true,
	}
	
	return s.generatedReportDAO.UpdateGeneratedReport(report)
}

// GetReportByChildIDAndType 获取指定类型的报告
func (s *AIReportService) GetReportByChildIDAndType(childArchiveID, reportType string) (*model.GeneratedReport, error) {
	return s.generatedReportDAO.GetGeneratedReportByChildIDAndType(childArchiveID, reportType)
}

// buildAIPrompt 构建AI请求提示词
func (s *AIReportService) buildAIPrompt(logs []model.HealingLog, reportType string) (string, error) {
	// 系统角色设定
	systemRole := `你是一位资深的儿童康复治疗师和心理健康专家，拥有超过15年的儿童发展和康复治疗经验。你专门从事儿童特殊需要康复、行为干预、情绪管理和发展评估工作。

你的专业背景包括：
- 儿童心理学和发展心理学专业知识
- 应用行为分析(ABA)治疗经验
- 感觉统合训练专业资质
- 家庭系统治疗和亲子关系指导经验
- 多元化康复方案设计和实施能力

请基于提供的疗愈记录数据，运用你的专业知识和临床经验，生成一份详细、专业且具有实际指导价值的分析报告。

`

	// 根据报告类型设定具体任务和示例结构
	var taskDescription string
	var analysisRequirements string
	var outputStructure string
	
	switch reportType {
	case "daily_summary":
		taskDescription = `**任务：生成日常疗愈总结报告**

你需要分析儿童在指定时间段内的疗愈进展，重点关注日常表现的变化和改善情况。这份报告将帮助家长和治疗团队了解儿童的当前状态和进步情况。

`
		analysisRequirements = `**详细分析要求：**

请从以下维度进行深入分析，每个维度都要提供具体的观察和评估：

1. **整体表现评估**（200-250字）
   - 总结儿童在此期间的整体康复态度和配合程度
   - 描述儿童的精神状态和活力水平变化
   - 评估治疗参与度和主动性表现
   - 分析整体发展趋势和康复进程

2. **进步亮点识别**（250-300字）
   - 详细描述3-4个最显著的进步表现，每个进步都要有具体例证
   - 量化改善程度（如：从每天3次情绪爆发减少到1次）
   - 对比治疗前后的具体变化
   - 突出里程碑式的突破和成就

3. **行为模式深度分析**（200-250字）
   - 分析儿童的日常行为规律和模式变化
   - 识别触发积极/消极行为的环境因素
   - 评估自我调节能力的发展情况
   - 描述适应性行为的增加和问题行为的减少

4. **社交互动能力评估**（150-200字）
   - 评估与治疗师、家长、同伴的互动质量
   - 分析沟通技巧和表达能力的变化
   - 描述社交主动性和合作能力的发展
   - 评估情绪共鸣和社会认知能力

5. **学习认知能力表现**（150-200字）
   - 分析注意力持续时间和集中度变化
   - 评估记忆力、理解力和执行功能
   - 描述新技能学习速度和掌握程度
   - 分析问题解决能力和创造性思维发展

6. **需要持续关注的挑战**（100-150字）
   - 客观识别仍需改进的具体方面
   - 分析可能的发展瓶颈和障碍
   - 提出需要加强的技能领域
   - 预警可能出现的退步风险

`
		outputStructure = `**报告结构示例：**
# 儿童疗愈总结报告

## 整体表现评估
[详细描述整体状态和康复态度]

## 进步亮点
- **[具体能力]显著提升**：[具体描述和例证]
- **[另一能力]明显改善**：[具体描述和例证]
[继续列出其他进步点]

## 行为模式分析
[分析行为规律和变化趋势]

## 社交互动情况
[评估社交能力发展]

## 学习能力表现
[分析认知和学习表现]

## 需要关注的问题
[客观指出需要改进的方面]

## 阶段性成果总结
[总结本阶段治疗成效和建议]

`

	case "suggestions":
		taskDescription = `**任务：生成康复建议报告**

你需要基于疗愈记录提供专业的康复建议和下一步治疗方案。这份报告将为治疗团队和家长提供具体可操作的指导建议。

`
		analysisRequirements = `**详细建议要求：**

请提供以下方面的专业建议，每个建议都要具体可操作：

1. **治疗方案优化建议**（300-400字）
   - 基于当前进展调整治疗重点和策略
   - 提供3-4个具体的治疗技术和方法
   - 建议治疗频率和强度调整
   - 制定个性化的干预计划

2. **日常生活技能发展计划**（250-300字）
   - 自理能力提升的具体训练方法
   - 运动协调能力发展建议
   - 生活技能学习的阶段性目标
   - 独立性培养的渐进式方案

3. **环境支持和优化建议**（200-250字）
   - 家庭环境结构化改善建议
   - 学校配合和支持方案
   - 社交环境优化措施
   - 感官环境调节建议

4. **长期发展规划**（200-250字）
   - 短期目标（1-3个月）的具体设定
   - 中期目标（3-6个月）的发展方向
   - 长期目标（6-12个月）的愿景规划
   - 各阶段的评估指标和里程碑

5. **家长参与和配合指导**（200-250字）
   - 家庭治疗配合的具体方法
   - 日常观察和记录的要点
   - 亲子互动技巧和策略
   - 家长心理支持和自我照顾建议

6. **风险防范和应急预案**（150-200字）
   - 识别潜在的治疗风险和挑战
   - 制定行为危机的应对策略
   - 建立支持网络和求助渠道
   - 定期评估和方案调整机制

`
		outputStructure = `**报告结构示例：**
# 儿童疗愈建议报告

## 治疗方案优化建议

### [具体治疗领域]强化
[详细的训练方法和技巧]

### [另一治疗领域]发展计划
[系统性的培养方案]

## 日常生活技能发展
### 自理能力提升
[具体的训练建议]

### 运动协调能力
[发展计划和活动建议]

## 环境支持建议
### 家庭环境优化
[具体的改善措施]

### 学校配合方案
[协作建议和支持方案]

## 长期发展规划
### 阶段性目标设定
[具体的时间节点和目标]

### 持续监测指标
[评估标准和观察要点]

## 家长参与建议
[详细的配合指导]

## 注意事项与风险防范
[风险识别和应对策略]

`

	case "progress":
		taskDescription = `**任务：生成进度分析报告**

你需要分析儿童的康复进度，评估治疗效果和发展趋势。这份报告将帮助评估当前治疗方案的有效性并指导后续调整。

`
		analysisRequirements = `**详细进度分析要求：**

请进行以下深度分析，尽可能量化和具体化：

1. **基线对比分析**（250-300字）
   - 与治疗初期状态进行详细对比
   - 量化各项能力的改善程度
   - 使用具体数据和百分比描述进步
   - 识别最显著的变化领域

2. **发展轨迹评估**（200-250字）
   - 分析各项能力的发展速度和趋势
   - 评估发展的稳定性和持续性
   - 识别发展的关键转折点
   - 预测未来发展方向

3. **治疗效果量化评估**（200-250字）
   - 评估不同治疗方法的有效性
   - 分析治疗目标的达成情况
   - 量化投入产出比和效率
   - 识别最有效的干预策略

4. **里程碑达成情况**（150-200字）
   - 评估重要发展里程碑的达成状况
   - 分析达成时间是否符合预期
   - 识别超预期和滞后的发展领域
   - 调整后续里程碑设定

5. **瓶颈和挑战分析**（150-200字）
   - 识别当前面临的主要发展瓶颈
   - 分析阻碍进步的内外因素
   - 评估挑战的严重程度和影响
   - 提出突破瓶颈的策略建议

6. **未来发展预测**（100-150字）
   - 基于当前趋势预测未来发展
   - 评估达到长期目标的可能性
   - 识别需要重点关注的发展领域
   - 建议治疗方案的调整方向

`
		outputStructure = `**报告结构示例：**
# 儿童康复进度分析报告

## 基线对比分析
[详细的前后对比和量化数据]

## 发展轨迹评估
[各能力发展趋势分析]

## 治疗效果量化评估
[具体的效果评估和数据]

## 里程碑达成情况
[重要节点的达成分析]

## 瓶颈和挑战分析
[发展障碍的识别和分析]

## 未来发展预测
[基于数据的发展预测]

`

	default:
		taskDescription = `**任务：生成综合分析报告**
请对儿童的整体疗愈情况进行全面综合分析。

`
		analysisRequirements = `**分析要求：**
请进行全面的综合分析，包括进展评估、问题识别和改进建议。

`
		outputStructure = `**报告结构：**
请按照专业报告格式组织内容。

`
	}
	
	// 构建疗愈记录数据部分
	var logsData strings.Builder
	logsData.WriteString("**疗愈记录数据分析：**\n\n")
	if len(logs) == 0 {
		logsData.WriteString("**数据状态：** 当前暂无具体的疗愈记录数据。\n\n")
		logsData.WriteString("**分析说明：** 由于缺乏具体的疗愈记录，请基于你的专业经验和临床知识，生成一份符合儿童康复治疗标准的专业报告。报告应体现典型的康复进展模式和专业的治疗建议。\n\n")
	} else {
		logsData.WriteString(fmt.Sprintf("**数据概览：** 共收集到 %d 条疗愈记录，时间跨度从 %s 到 %s\n\n", 
			len(logs), 
			logs[0].CreatedAt.Format("2006年01月02日"), 
			logs[len(logs)-1].CreatedAt.Format("2006年01月02日")))
		
		logsData.WriteString("**详细记录内容：**\n")
		for i, log := range logs {
			logsData.WriteString(fmt.Sprintf("\n**记录 %d**（%s）\n", i+1, log.CreatedAt.Format("2006年01月02日 15:04")))
			logsData.WriteString(fmt.Sprintf("- **记录内容：** %s\n", log.Content))
			
			// 如果有媒体文件，详细描述
			if len(log.Media) > 0 {
				logsData.WriteString("- **附件媒体：** ")
				var mediaDetails []string
				for _, media := range log.Media {
					mediaDetails = append(mediaDetails, fmt.Sprintf("%s文件", media.MediaType))
				}
				logsData.WriteString(strings.Join(mediaDetails, "、"))
				logsData.WriteString("（这些媒体文件提供了额外的行为观察和进展证据）\n")
			}
			
			// 添加记录分析提示
			logsData.WriteString("- **分析要点：** 请重点关注此记录中体现的行为变化、情绪状态、技能表现和社交互动情况\n")
		}
		logsData.WriteString("\n**数据分析指导：** 请基于以上记录内容，结合时间序列分析儿童的发展变化趋势，识别进步模式和需要关注的问题。\n\n")
	}
	
	// 专业要求和质量标准
	professionalRequirements := `**专业标准和质量要求：**

1. **专业性要求：**
   - 使用儿童康复治疗领域的专业术语和概念
   - 体现循证实践的理念和方法
   - 展现多学科整合的治疗视角
   - 遵循儿童发展的科学规律

2. **内容深度要求：**
   - 每个分析点都要有具体的观察描述和专业解释
   - 提供可量化的改善指标和具体例证
   - 包含对行为背后原因的深层分析
   - 给出基于证据的专业判断和建议

3. **实用性要求：**
   - 所有建议都必须具体可操作
   - 提供明确的实施步骤和方法
   - 考虑家庭和学校的实际执行条件
   - 包含风险评估和应对策略

4. **语言表达要求：**
   - 使用温暖、积极但客观的专业语调
   - 避免过于技术性的术语，确保家长能够理解
   - 平衡希望与现实，既要鼓励又要客观
   - 体现对儿童和家庭的尊重与支持

`
	
	// 输出格式要求
	formatRequirements := `**输出格式和结构要求：**

1. **格式规范：**
   - 严格使用Markdown格式
   - 使用清晰的标题层级（#、##、###）
   - 合理使用列表、加粗、斜体等格式
   - 确保排版美观、层次分明

2. **内容长度：**
   - 总报告长度控制在1200-1500字
   - 各部分内容要均衡分配
   - 重点部分可以适当详细
   - 避免冗余和重复表述

3. **结构完整性：**
   - 必须包含所有要求的分析维度
   - 每个部分都要有实质性内容
   - 逻辑清晰，前后呼应
   - 结论明确，建议具体

4. **专业报告标准：**
   - 开头要有简明的概述
   - 中间部分要有详细的分析
   - 结尾要有明确的总结和建议
   - 整体体现专业水准和临床价值

**现在请开始生成专业的儿童康复分析报告：**

`
	
	// 组合所有部分
	prompt := systemRole + taskDescription + logsData.String() + analysisRequirements + outputStructure + professionalRequirements + formatRequirements
	
	return prompt, nil
}

// callAIAPI 调用AI API生成内容
func (s *AIReportService) callAIAPI(prompt string) (string, error) {
	// 获取AI配置
	aiConfig := config.GetAIConfig()
	
	// 检查是否配置了API密钥
	if aiConfig.APIKey == "" || aiConfig.APIKey == "your_ai_api_key_here" {
		// 如果没有配置真实的API密钥，使用模拟数据
		return s.generateMockContent(prompt)
	}
	
	// 使用真实的AI API
	return s.callRealAIAPI(prompt)
}

// generateMockContent 生成模拟内容
func (s *AIReportService) generateMockContent(prompt string) (string, error) {
	if strings.Contains(prompt, "summary") {
		return s.generateMockSummary(), nil
	} else if strings.Contains(prompt, "suggestion") {
		return s.generateMockSuggestion(), nil
	}
	return "AI生成的模拟内容", nil
}

// generateMockSummary 生成模拟总结报告
func (s *AIReportService) generateMockSummary() string {
	return `# 儿童疗愈总结报告

## 整体表现评估
在本阶段的疗愈过程中，儿童整体表现良好，展现出积极的康复态度和明显的进步迹象。通过系统性的治疗干预，儿童在多个发展维度上都取得了可观的改善。

## 进步亮点
- **情绪调节能力显著提升**：从初期的情绪波动较大，到现在能够较好地控制和表达情绪
- **社交互动意愿增强**：主动与治疗师和同伴交流的频率明显增加
- **注意力集中度改善**：能够在活动中保持更长时间的专注
- **自理能力提升**：在日常生活技能方面表现出更强的独立性

## 行为模式分析
儿童的行为模式呈现出积极的变化趋势：
- 晨间情绪状态更加稳定
- 对新环境和新活动的适应能力增强
- 规则意识和配合度显著提高
- 挫折耐受性有所改善

## 社交互动情况
- 与治疗师建立了良好的信任关系
- 开始主动寻求帮助和支持
- 在小组活动中表现出更好的合作精神
- 对他人情绪的感知和回应能力提升

## 学习能力表现
- 新技能学习速度加快
- 记忆保持能力增强
- 模仿学习能力显著提升
- 对指令的理解和执行更加准确

## 需要关注的问题
- 在面对复杂任务时仍需要更多支持
- 某些特定情境下的焦虑反应需要继续关注
- 精细动作技能还有进一步提升空间

## 阶段性成果总结
本阶段治疗取得了显著成效，儿童在情绪管理、社交互动、学习能力等方面都有明显进步。建议继续当前的治疗方案，并适当增加挑战性活动以促进进一步发展。`
}

// generateMockSuggestion 生成模拟建议报告
func (s *AIReportService) generateMockSuggestion() string {
	return `# 儿童疗愈建议报告

## 治疗方案优化建议

### 情绪管理训练强化
基于当前观察到的情绪调节能力提升，建议进一步深化情绪管理训练：
- **情绪识别练习**：使用情绪卡片和表情镜像游戏，帮助儿童更准确地识别和命名情绪
- **情绪调节技巧**：教授深呼吸、数数法、正念冥想等具体的情绪调节策略
- **情绪表达训练**：通过角色扮演和情境模拟，练习适当的情绪表达方式

### 社交技能发展计划
针对社交互动能力的改善趋势，制定系统性的社交技能培养方案：
- **同伴互动活动**：组织小组游戏和合作任务，增加与同龄人的互动机会
- **沟通技巧训练**：练习主动打招呼、请求帮助、表达需求等基础社交技能
- **冲突解决能力**：教授处理分歧和冲突的健康方式

### 认知能力提升策略
基于学习能力的积极表现，建议实施以下认知训练：
- **注意力训练**：通过专注力游戏和任务，进一步提升持续注意力
- **记忆力强化**：使用记忆游戏和重复练习，巩固记忆保持能力
- **问题解决训练**：设计适龄的逻辑推理和问题解决任务

## 日常生活技能发展

### 自理能力提升
- **生活技能训练**：系统性地教授和练习日常生活必需技能
- **独立性培养**：逐步减少辅助，鼓励独立完成适龄任务
- **责任感建立**：分配适当的家务和责任，培养责任意识

### 运动协调能力
- **大运动技能**：通过体育活动和户外游戏提升大肌肉群协调性
- **精细动作训练**：使用手工制作、绘画等活动提升手眼协调能力
- **感觉统合训练**：针对感觉处理问题进行专项训练

## 环境支持建议

### 家庭环境优化
- **结构化环境**：建立清晰的日常作息和规则体系
- **积极强化**：及时给予正面反馈和鼓励，增强自信心
- **安全感建立**：创造稳定、可预测的家庭环境

### 学校配合方案
- **个别化教育计划**：与学校合作制定适合的学习支持方案
- **教师沟通**：定期与教师交流，确保治疗目标的一致性
- **同伴支持**：培养同学间的理解和支持关系

## 长期发展规划

### 阶段性目标设定
- **短期目标（1-3个月）**：巩固当前进步，重点提升情绪稳定性
- **中期目标（3-6个月）**：扩展社交圈子，提升学习适应能力
- **长期目标（6-12个月）**：实现更高程度的独立性和社会适应性

### 持续监测指标
- 情绪调节频率和效果
- 社交互动的主动性和质量
- 学习任务的完成情况
- 日常生活技能的掌握程度

## 家长参与建议

### 家庭治疗配合
- **一致性原则**：确保家庭和治疗环境的方法一致
- **耐心支持**：给予充分的时间和空间让儿童适应和成长
- **积极参与**：主动学习相关知识，参与治疗过程

### 日常观察记录
- 记录儿童的行为变化和进步表现
- 注意观察可能的退步信号
- 及时与治疗团队沟通反馈

## 注意事项与风险防范

### 潜在风险识别
- 避免过度期待导致的压力
- 防止治疗疲劳和抗拒情绪
- 注意个体差异，避免一刀切的方法

### 应急处理预案
- 制定情绪爆发时的应对策略
- 建立支持网络和求助渠道
- 定期评估和调整治疗方案

通过以上综合性的建议和支持措施，相信能够进一步促进儿童的全面发展和康复进程。`
}

// 实际项目中的AI API调用示例（需要根据具体的AI服务进行调整）
func (s *AIReportService) callRealAIAPI(prompt string) (string, error) {
	// 获取AI配置
	aiConfig := config.GetAIConfig()
	
	// 构建请求体 - 使用系统消息和用户消息的组合
	requestBody := map[string]interface{}{
		"model": aiConfig.Model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "你是一位专业的儿童康复治疗师和心理健康专家，拥有丰富的儿童发展和康复治疗经验。你需要基于提供的疗愈记录数据，生成专业、详细且具有指导意义的分析报告。请确保报告内容专业准确，语言清晰易懂，建议具体可操作。",
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"max_tokens":  aiConfig.MaxTokens,
		"temperature": aiConfig.Temperature,
	}
	
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("构建请求体失败: %v", err)
	}
	
	// 创建HTTP请求
	req, err := http.NewRequest("POST", aiConfig.BaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+aiConfig.APIKey)
	
	// 发送请求
	client := &http.Client{
		Timeout: time.Duration(aiConfig.Timeout) * time.Second,
	}
	
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()
	
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}
	
	// 解析响应
	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error struct {
			Message string `json:"message"`
			Type    string `json:"type"`
		} `json:"error"`
	}
	
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}
	
	// 检查是否有错误
	if response.Error.Message != "" {
		return "", fmt.Errorf("AI API错误: %s (%s)", response.Error.Message, response.Error.Type)
	}
	
	if len(response.Choices) == 0 {
		return "", fmt.Errorf("AI响应为空")
	}
	
	content := response.Choices[0].Message.Content
	if content == "" {
		return "", fmt.Errorf("AI生成的内容为空")
	}
	
	return content, nil
}