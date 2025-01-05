package mock

import (
	"github.com/babycare/models"
	"gorm.io/gorm"
)

func GenerateMockData(db *gorm.DB) error {
	// 创建用户
	user := models.User{
		Username: "专家顾问",
		Email:    "expert@babycare.com",
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// 孕期文章数据
	pregnancyArticles := []models.Post{
		{
			Title:     "孕期营养食谱指南",
			Content:   "孕期营养对胎儿的发育至关重要。本文将为您详细介绍孕期各阶段的营养需求和推荐食谱。\n\n第一孕期(1-3个月)：\n- 叶酸rich食物：绿叶蔬菜、豆类\n- 优质蛋白：瘦肉、鱼、蛋\n- 注意事项：避免生食、酒精\n\n第二孕期(4-6个月)：\n- 补充钙质：牛奶、酸奶、虾皮\n- 铁质补充：红肉、菠菜\n- 推荐食谱：牛奶燕麦、菠菜炒蛋\n\n第三孕期(7-9个月)：\n- 控制总热量\n- 适量优质蛋白\n- 注意补钙补铁",
			UserID:    user.ID,
			Tags:      "营养,食谱,孕期指导",
			Views:     100,
			WordCount: 500,
			ReadTime:  5,
			Category:  "pregnancy",
		},
		{
			Title:     "孕期营养补充剂使用指南",
			Content:   "科学使用营养补充剂是保障孕期营养的重要手段。本文将详细介绍各类营养补充剂的作用和使用方法。\n\n1. 叶酸\n- 服用时间：孕前3个月到孕后3个月\n- 推荐剂量：400-800微克/天\n- 注意事项：空腹服用\n\n2. 铁剂\n- 服用时间：孕中期开始\n- 推荐剂量：视血红蛋白水平而定\n- 注意事项：与维生素C同服\n\n3. 钙片\n- 服用时间：全孕期\n- 推荐剂量：1000mg/天\n- 注意事项：避免与铁剂同服",
			UserID:    user.ID,
			Tags:      "营养,补充剂,孕期保健",
			Views:     80,
			WordCount: 400,
			ReadTime:  4,
			Category:  "pregnancy",
		},
		{
			Title:     "孕期瑜伽入门指南",
			Content:   "孕期瑜伽可以帮助准妈妈保持身体柔韧性，缓解不适，为分娩做准备。\n\n安全注意事项：\n1. 选择专业的孕期瑜伽教练\n2. 避免过度拉伸\n3. 注意呼吸节奏\n\n推荐动作：\n1. 猫牛式\n- 缓解背部疼痛\n- 增强核心力量\n\n2. 蝴蝶式\n- 打开髋部\n- 改善血液循环\n\n3. 深呼吸练习\n- 放松身心\n- 增加肺活量",
			UserID:    user.ID,
			Tags:      "运动,瑜伽,孕期锻炼",
			Views:     120,
			WordCount: 300,
			ReadTime:  3,
			Category:  "pregnancy",
		},
		{
			Title:     "孕期散步指南",
			Content:   "散步是最安全和推荐的孕期运动之一。本文将介绍孕期散步的益处和注意事项。\n\n散步益处：\n1. 改善血液循环\n2. 预防妊娠糖尿病\n3. 缓解腿部浮肿\n\n建议时间：\n- 早晚温度适宜时\n- 每次20-30分钟\n- 每周3-5次\n\n注意事项：\n1. 选择平坦的路面\n2. 穿着舒适的鞋子\n3. 注意天气变化",
			UserID:    user.ID,
			Tags:      "运动,散步,户外活动",
			Views:     90,
			WordCount: 350,
			ReadTime:  4,
			Category:  "pregnancy",
		},
		{
			Title:     "孕期游泳指南",
			Content:   "游泳是非常适合孕妇的运动，可以全面锻炼身体，减轻水中重力对身体的压力。\n\n游泳益处：\n1. 减轻关节压力\n2. 改善心肺功能\n3. 缓解背部不适\n\n注意事项：\n1. 选择合适的泳池\n- 水温适中\n- 卫生条件好\n\n2. 运动强度\n- 以轻度为主\n- 避免剧烈动作\n\n3. 安全措施\n- 做好热身\n- 随身携带毛巾",
			UserID:    user.ID,
			Tags:      "运动,游泳,水中运动",
			Views:     70,
			WordCount: 320,
			ReadTime:  3,
			Category:  "pregnancy",
		},
	}

	// 批量创建文章
	if err := db.Create(&pregnancyArticles).Error; err != nil {
		return err
	}

	// 新生儿护理文章数据
	newbornArticles := []models.Post{
		{
			Title:     "新生儿母乳喂养指南",
			Content:   "母乳喂养是新生儿最理想的喂养方式。本文将为您详细介绍母乳喂养的要点。\n\n正确姿势：\n1. 抱姿要求\n- 婴儿头、颈、身体保持一条直线\n- 腹贴腹，面对面\n- 注意支撑婴儿头部\n\n喂养时间：\n- 新生儿每2-3小时喂一次\n- 持续时间15-20分钟\n- 按需喂养更科学\n\n注意事项：\n1. 保持乳房卫生\n2. 确保吸吮正确\n3. 避免乳汁淤积",
			UserID:    user.ID,
			Tags:      "喂养,母乳,新生儿护理",
			Views:     150,
			WordCount: 400,
			ReadTime:  4,
			Category:  "newborn",
		},
		{
			Title:     "新生儿人工喂养建议",
			Content:   "当母乳喂养不足或无法进行时，科学的人工喂养同样重要。\n\n奶粉选择：\n1. 选择适合月龄\n2. 注意品牌可靠性\n3. 适度调整配方\n\n冲调方法：\n- 使用40-50度温水\n- 严格按比例冲调\n- 摇匀避免结块\n\n喂养技巧：\n1. 选择合适奶嘴\n2. 控制奶温\n3. 及时消毒",
			UserID:    user.ID,
			Tags:      "喂养,奶粉,新生儿护理",
			Views:     120,
			WordCount: 350,
			ReadTime:  3,
			Category:  "newborn",
		},
		{
			Title:     "新生儿脐带护理指南",
			Content:   "正确的脐带护理对预防感染至关重要。\n\n基本步骤：\n1. 保持干燥\n- 每天观察\n- 避免沾水\n\n2. 消毒方法\n- 使用医用酒精\n- 棉签蘸取适量\n- 由内向外擦拭\n\n3. 注意事项\n- 勿包裹过紧\n- 避免使用药粉\n- 及时更换尿布",
			UserID:    user.ID,
			Tags:      "护理,脐带,新生儿护理",
			Views:     200,
			WordCount: 300,
			ReadTime:  3,
			Category:  "newborn",
		},
		{
			Title:     "新生儿洗澡全攻略",
			Content:   "正确的洗澡方式能让宝宝舒适健康。\n\n准备工作：\n1. 室温保持在26度以上\n2. 水温37-38度\n3. 准备好所需用品\n\n洗澡步骤：\n1. 先洗头部和脸部\n2. 依次清洗颈部、上肢\n3. 清洗躯干和下肢\n4. 最后清洗臀部\n\n注意事项：\n1. 动作轻柔\n2. 时间控制在5-10分钟\n3. 及时擦干保暖",
			UserID:    user.ID,
			Tags:      "护理,洗澡,新生儿护理",
			Views:     180,
			WordCount: 380,
			ReadTime:  4,
			Category:  "newborn",
		},
		{
			Title:     "新生儿换尿布指南",
			Content:   "科学的换尿布方法可以预防红臀。\n\n更换时机：\n1. 大小便后及时更换\n2. 定时检查尿布\n3. 夜间适当延长间隔\n\n操作步骤：\n1. 准备清洁用品\n2. 轻柔擦拭方向\n3. 充分清洁后擦干\n\n预防红臀：\n1. 使用护臀霜\n2. 保持局部干燥\n3. 适当透气时间",
			UserID:    user.ID,
			Tags:      "护理,尿布,新生儿护理",
			Views:     160,
			WordCount: 320,
			ReadTime:  3,
			Category:  "newborn",
		},
	}

	// 批量创建新生儿文章
	if err := db.Create(&newbornArticles).Error; err != nil {
		return err
	}

	// 早期发展文章数据
	developmentArticles := []models.Post{
		{
			Title:     "新生儿视觉发展指南",
			Content:   "了解新生儿视觉发展的关键阶段。\n\n发展阶段：\n1. 出生后1个月\n- 能看清20-30厘米内的物体\n- 对强烈的光影反应敏感\n- 喜欢黑白对比的图案\n\n2. 2-3个月\n- 开始对色彩产生兴趣\n- 能跟随移动物体\n- 认识熟悉的面孔\n\n视觉训练：\n1. 选择合适玩具\n- 黑白对比图卡\n- 带有鲜艳色彩的玩具\n- 会动的悬挂玩具\n\n2. 互动方式\n- 经常变换婴儿视角\n- 适当的面对面交流\n- 控制观察时间",
			UserID:    user.ID,
			Tags:      "发展,视觉,新生儿护理",
			Views:     90,
			WordCount: 350,
			ReadTime:  4,
			Category:  "newborn",
		},
		{
			Title:     "婴儿听觉训练方法",
			Content:   "科学的听觉训练对婴儿语言发展至关重要。\n\n训练方法：\n1. 音乐互动\n- 选择舒缓的音乐\n- 控制音量大小\n- 配合律动\n\n2. 语言交流\n- 多与宝宝说话\n- 模仿宝宝发音\n- 简单的音节练习\n\n3. 环境音识别\n- 介绍日常声音\n- 培养声音方位感\n- 声音大小辨别",
			UserID:    user.ID,
			Tags:      "发展,听觉,新生儿护理",
			Views:     85,
			WordCount: 300,
			ReadTime:  3,
			Category:  "newborn",
		},
		{
			Title:     "婴儿早期运动技能发展",
			Content:   "支持婴儿运动技能发展的科学方法。\n\n关键里程碑：\n1. 抬头\n- 俯卧时间训练\n- 颈部肌肉锻炼\n- 安全注意事项\n\n2. 翻身\n- 适当的引导方式\n- 训练时机选择\n- 保护措施\n\n3. 坐立\n- 循序渐进训练\n- 核心力量培养\n- 辅助工具使用",
			UserID:    user.ID,
			Tags:      "发展,运动,新生儿护理",
			Views:     95,
			WordCount: 320,
			ReadTime:  3,
			Category:  "newborn",
		},
		{
			Title:     "建立婴儿健康作息指南",
			Content:   "科学的作息规律对婴儿成长至关重要。\n\n作息安排：\n1. 睡眠规律\n- 认识睡眠规律\n- 建立睡眠仪式\n- 创造良好环境\n\n2. 喂养时间\n- 根据月龄调整\n- 避免夜间频繁喂养\n- 建立进食规律\n\n3. 活动安排\n- 适量户外活动\n- 互动游戏时间\n- 避免过度刺激",
			UserID:    user.ID,
			Tags:      "发展,作息,新生儿护理",
			Views:     88,
			WordCount: 330,
			ReadTime:  3,
			Category:  "newborn",
		},
	}

	// 批量创建早期发展文章
	if err := db.Create(&developmentArticles).Error; err != nil {
		return err
	}

	return nil
} 