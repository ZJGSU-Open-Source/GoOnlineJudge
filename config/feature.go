package config

const (
	ProblemPerPage  = 100
	ContestPerPage  = 100
	ExercisePerPage = 100
	SolutionPerPage = 50
	UserPerPage     = 50
)

const (
	PageHeadLimit = 1
	PageTailLimit = 1
	PageMidLimit  = 2
)

const (
	JudgeNA  = 0  //None
	JudgePD  = 1  //Pending
	JudgeRJ  = 2  //Running & judgingconst
	JudgeAC  = 3  //Accepted
	JudgeCE  = 4  //Compile Error
	JudgeRE  = 5  //Runtime Error
	JudgeWA  = 6  //Wrong Answer
	JudgeTLE = 7  //Time Limit Exceeded
	JudgeMLE = 8  //Memory Limit Exceeded
	JudgeOLE = 9  //Output Limit Exceeded
	JudgeCP  = 10 //wait to Compare Output
)

const (
	LanguageNA   = 0 //None
	LanguageC    = 1 //C
	LanguageCPP  = 2 //C++
	LanguageJAVA = 3 //Java
)

const (
	ModuleNA = 0 //None
	ModuleP  = 1 //None
	ModuleC  = 2 //Contest
	ModuleE  = 3 //Exercise
)

const (
	PrivilegeNA = 0 //None
	PrivilegePU = 1 //Primary User
	PrivilegeSB = 2 //Source Broswer
	PrivilegeAD = 3 //Admin
)

const (
	EncryptNA = 0 //None
	EncryptPB = 1 //Public
	EncryptPT = 2 //Private
	EncryptPW = 3 //Password
)

const (
	StatusReverse   = 0 //不可用
	StatusIncon     = 1 //正在比赛中
	StatusAvailable = 2 //可用
	StatusPending   = 3 //等待
	StatusRunning   = 4 //进行中
	StatusEnding    = 5 //结束
)
