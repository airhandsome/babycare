package mock

import (
	"log"
	"time"

	"github.com/babycare/models"
	"gorm.io/gorm"
)

func GenerateMockData(db *gorm.DB) error {
	//return nil
	// 创建用户
	user := models.User{
		Username: "专家顾问",
		Email:    "expert@babycare.com",
		Password: "123456",
	}

	// 先查找用户是否存在
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		// 用户不存在，创建新用户
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to create user: %v", err)
			return err
		}
	} else {
		// 用户已存在，使用已存在的用户
		user = existingUser
	}

	// 专家专栏数据
	expertArticles := []models.Post{
		{
			Title:     "儿童心理专家张医生",
			Summary:   "从事儿童心理咨询15年，专注于儿童心理健康和行为发展",
			Content:   "专业背景与经验...",
			UserID:    user.ID,
			Tags:      "儿童心理,行为发展,情绪管理",
			Category:  "expert_profile",
			Avatar:    "https://images.unsplash.com/photo-1559839734-2b71ea197ec2",
			Views:     150,
			WordCount: 300,
			ReadTime:  3,
		},
		{
			Title:     "育儿经验分享：如何培养孩子的专注力",
			Summary:   "通过科学的方法和日常练习提升孩子的注意力集中能力",
			Content:   "详细的专注力培养方法...",
			UserID:    user.ID,
			Tags:      "专注力,教育方法,育儿经验",
			Category:  "expert_article",
			Views:     200,
			WordCount: 800,
			ReadTime:  8,
		},
		{
			Title:     "热门话题：儿童教育中的奖惩问题",
			Summary:   "探讨在教育过程中如何正确使用奖励和惩罚",
			Content:   "关于奖惩的深入讨论...",
			UserID:    user.ID,
			Tags:      "教育方法,奖惩制度,家庭教育",
			Category:  "expert_topic",
			Views:     180,
			WordCount: 600,
			ReadTime:  6,
		},
	}

	// 批量创建专家专栏文章
	if err := db.Create(&expertArticles).Error; err != nil {
		log.Printf("Failed to create expert articles: %v", err)
		return err
	}

	// 检查是否已经存在文章数据
	var count int64
	if err := db.Model(&models.Post{}).Count(&count).Error; err != nil {
		return err
	}

	// 如果已经存在数据，则跳过创建
	if count > 0 {
		log.Println("Posts data already exists, skipping creation")
		return nil
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
		log.Printf("Failed to create pregnancy articles: %v", err)
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
		log.Print(err)
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
		log.Printf("Failed to create development articles: %v", err)
		return err
	}

	// 创建论坛帖子
	forumPosts := []models.ForumPost{
		{
			UserID:    user.ID,
			Category:  "育儿经验",
			Content:   "分享一下如何让宝宝爱上吃蔬菜的小技巧！我家宝宝以前特别挑食，通过这些方法现在已经能主动吃蔬菜了...",
			Images:    "/images/forum/vegetables.png",
			Likes:     15,
			Comments:  8,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Category:  "专业指导",
			Content:   "关于婴儿睡眠的一些专业建议：1. 建立规律的作息时间...",
			Likes:     32,
			Comments:  12,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(&forumPosts).Error; err != nil {
		log.Printf("Failed to create forum posts: %v", err)
		return err
	}

	// 创建一些评论
	comments := []models.Comment{
		{
			UserID:    user.ID,
			PostID:    forumPosts[0].ID,
			Content:   "这个方法很有用，我家宝宝也是这样慢慢学会吃蔬菜的！",
			Likes:     5,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			PostID:    forumPosts[0].ID,
			Content:   "可以分享一下具体是怎么做的吗？",
			Likes:     3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(&comments).Error; err != nil {
		log.Printf("Failed to create comments: %v", err)
		return err
	}

	// 创建里程碑数据
	milestones := []models.Milestone{
		{
			UserID:      user.ID,
			Category:    "0-3个月",
			Title:       "社交微笑",
			Completed:   true,
			CompletedAt: time.Now().Add(-24 * time.Hour),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			UserID:      user.ID,
			Category:    "0-3个月",
			Title:       "抬头",
			Completed:   true,
			CompletedAt: time.Now().Add(-48 * time.Hour),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			UserID:    user.ID,
			Category:  "0-3个月",
			Title:     "追视物体",
			Completed: false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Category:  "4-6个月",
			Title:     "翻身",
			Completed: false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Category:  "4-6个月",
			Title:     "坐立",
			Completed: false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:      user.ID,
			Category:    "4-6个月",
			Title:       "发出声音",
			Completed:   true,
			CompletedAt: time.Now().Add(-12 * time.Hour),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := db.Create(&milestones).Error; err != nil {
		log.Printf("Failed to create milestones: %v", err)
		return err
	}

	// 创建成长记录数据
	growthRecords := []models.GrowthRecord{
		{
			UserID:    user.ID,
			Type:      "身高",
			Value:     50.5,
			Unit:      "cm",
			Note:      "出生身高",
			Date:      time.Now().Add(-90 * 24 * time.Hour),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Type:      "体重",
			Value:     3.5,
			Unit:      "kg",
			Note:      "出生体重",
			Date:      time.Now().Add(-90 * 24 * time.Hour),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Type:      "身高",
			Value:     55.0,
			Unit:      "cm",
			Note:      "一月例行检查",
			Date:      time.Now().Add(-60 * 24 * time.Hour),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Type:      "体重",
			Value:     4.2,
			Unit:      "kg",
			Note:      "一月例行检查",
			Date:      time.Now().Add(-60 * 24 * time.Hour),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Type:      "身高",
			Value:     58.5,
			Unit:      "cm",
			Note:      "两月例行检查",
			Date:      time.Now().Add(-30 * 24 * time.Hour),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    user.ID,
			Type:      "体重",
			Value:     5.1,
			Unit:      "kg",
			Note:      "两月例行检查",
			Date:      time.Now().Add(-30 * 24 * time.Hour),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := db.Create(&growthRecords).Error; err != nil {
		log.Printf("Failed to create growth records: %v", err)
		return err
	}

	// 幼儿教育文章数据
	educationArticles := []models.Post{
		{
			Title:     "选择合适的幼儿园指南",
			Content:   "如何为孩子选择合适的幼儿园环境？\n\n考虑因素：\n1. 教育理念\n- 是否符合家庭价值观\n- 教学方式是否适合孩子\n- 师资力量评估\n\n2. 环境设施\n- 安全措施是否完善\n- 活动空间是否充足\n- 卫生状况评估\n\n3. 地理位置\n- 距离家庭远近\n- 交通便利程度\n- 周边环境评估",
			UserID:    user.ID,
			Tags:      "教育指南,环境选择,幼儿园",
			Views:     120,
			WordCount: 400,
			ReadTime:  5,
			Category:  "education_guide",
			Image:     "https://images.unsplash.com/photo-1503454537195-1dcabb73ffb9",
		},
		{
			Title:     "幼儿社交能力培养",
			Content:   "培养孩子良好的社交能力对未来发展至关重要。\n\n培养方法：\n1. 角色扮演游戏\n- 模拟日常社交场景\n- 培养换位思考能力\n- 提升表达能力\n\n2. 团体活动参与\n- 组织小组游戏\n- 培养合作意识\n- 学习分享行为",
			UserID:    user.ID,
			Tags:      "社交技能,教育指导,幼儿发展",
			Views:     95,
			WordCount: 350,
			ReadTime:  4,
			Category:  "social_skills",
			Icon:      "Share",
		},
		{
			Title:     "创意美术活动设计",
			Content:   "通过艺术活动激发孩子的创造力和想象力。\n\n活动设计：\n1. 绘画探索\n- 不同材料尝试\n- 主题创作引导\n- 自由表达鼓励\n\n2. 手工制作\n- 材料安全选择\n- 难度适当把控\n- 成就感培养",
			UserID:    user.ID,
			Tags:      "艺术创造,美术活动,幼儿教育",
			Views:     88,
			WordCount: 320,
			ReadTime:  4,
			Category:  "art_activities",
			Image:     "https://images.unsplash.com/photo-1513364776144-60967b0f800f",
			Duration:  "30分钟",
			Materials: "画纸,蜡笔,水彩",
		},
		// 社交技能文章
		{
			Title:     "分享与合作能力培养",
			Summary:   "通过日常活动和游戏培养孩子的分享意识和团队合作精神",
			Content:   "详细介绍培养孩子分享与合作能力的方法...\n\n1. 日常培养\n- 从简单物品分享开始\n- 表扬分享行为\n- 以身作则\n\n2. 游戏方式\n- 团队合作游戏\n- 角色扮演\n- 集体创作活动\n\n3. 情境教学\n- 创造分享机会\n- 处理冲突指导\n- 正面引导方式",
			UserID:    user.ID,
			Tags:      "社交技能,分享,合作",
			Category:  "social_skills",
			Icon:      "Share",
			Views:     85,
			WordCount: 500,
			ReadTime:  6,
		},
		{
			Title:     "情绪管理技巧",
			Summary:   "帮助孩子认识和管理自己的情绪，建立健康的情绪表达方式",
			Content:   "儿童情绪管理的重要性和具体方法...\n\n1. 情绪认知\n- 识别基本情绪\n- 理解情绪原因\n- 接纳情绪变化\n\n2. 管理技巧\n- 情绪表达方式\n- 自我安抚方法\n- 寻求帮助意识\n\n3. 家长指导\n- 情绪共情\n- 行为引导\n- 建立规则",
			UserID:    user.ID,
			Tags:      "情绪管理,社交技能,心理健康",
			Category:  "social_skills",
			Icon:      "Sunny",
			Views:     92,
			WordCount: 480,
			ReadTime:  5,
		},
		// 艺术活动
		{
			Title:     "音乐律动课程",
			Summary:   "通过音乐和肢体活动培养儿童艺术感知和表达能力",
			Content:   "详细的音乐律动课程设计和活动方案...\n\n1. 课程设计\n- 年龄适配内容\n- 循序渐进安排\n- 多元化活动\n\n2. 活动方式\n- 歌曲学习\n- 乐器体验\n- 舞蹈创作\n\n3. 教具准备\n- 乐器选择\n- 音乐资源\n- 场地布置",
			UserID:    user.ID,
			Tags:      "音乐,艺术,律动",
			Category:  "art_activities",
			Image:     "https://images.unsplash.com/photo-1445633629932-0029acc44e88",
			Duration:  "45分钟",
			Materials: "乐器,音乐播放器,节奏卡片",
			Views:     78,
			WordCount: 450,
			ReadTime:  5,
		},
		{
			Title:     "手工艺术探索",
			Summary:   "激发创造力的手工艺术活动，培养精细动作和艺术表现力",
			Content:   "丰富多样的手工艺术活动设计...\n\n1. 材料运用\n- 安全材料选择\n- 工具使用指导\n- 创新材料搭配\n\n2. 项目设计\n- 趣味性主题\n- 难度递进\n- 成果展示\n\n3. 能力培养\n- 空间认知\n- 色彩搭配\n- 创意表达",
			UserID:    user.ID,
			Tags:      "手工,艺术,创造力",
			Category:  "art_activities",
			Image:     "https://images.unsplash.com/photo-1513364776144-60967b0f800f",
			Duration:  "40分钟",
			Materials: "彩纸,剪刀,胶水,装饰物",
			Views:     82,
			WordCount: 420,
			ReadTime:  5,
		},
	}

	// 批量创建幼儿教育文章
	if err := db.Create(&educationArticles).Error; err != nil {
		log.Printf("Failed to create education articles: %v", err)
		return err
	}

	return nil
}
