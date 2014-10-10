package config

//server
const CookieExpires = 1800
const Datapath = "../ProblemData/"
const JudgeHost = "http://127.0.0.1:8888"

//CONSTANT
const (
	ProblemPerPage  = 2
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
	JudgePD  = 0  //Pending
	JudgeRJ  = 1  //Running & judging
	JudgeCE  = 2  //Compile Error
	JudgeAC  = 3  //Accepted
	JudgeRE  = 4  //Runtime Error
	JudgeWA  = 5  //Wrong Answer
	JudgeTLE = 6  //Time Limit Exceeded
	JudgeMLE = 7  //Memory Limit Exceeded
	JudgeOLE = 8  //Output Limit Exceeded
	JudgePE  = 9  //Presentation Error
	JudgeNA  = 10 //System Error
	JudgeRPD = 11 //Rejudge Pending
)

const (
	LanguageNA   = 0 //None
	LanguageC    = 1 //C
	LanguageCPP  = 2 //C++
	LanguageJAVA = 3 //Java
)

const (
	ModuleNA = 0 //None
	ModuleP  = 1 //Problem
	ModuleC  = 2 //Contest
	ModuleE  = 3 //Exercise
)

const (
	PrivilegeNA = 0 //None
	PrivilegePU = 1 //Primary User
	PrivilegeTC = 2 //Teacher
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
