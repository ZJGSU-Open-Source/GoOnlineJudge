package config

const ProblemPerPage = 100
const ContestPerPage = 100
const ExercisePerPage = 100
const SolutionPerPage = 50
const UserPerPage = 50

// const ProblemPerPage = 1
// const ContestPerPage = 1
// const ExercisePerPage = 1
// const SolutionPerPage = 5
// const UserPerPage = 50

const PageHeadLimit = 1
const PageTailLimit = 1
const PageMidLimit = 2

const JudgeNA = 0  //None
const JudgePD = 1  //Pending
const JudgeRJ = 2  //Running & judging
const JudgeAC = 3  //Accepted
const JudgeCE = 4  //Compile Error
const JudgeRE = 5  //Runtime Error
const JudgeWA = 6  //Wrong Answer
const JudgeTLE = 7 //Time Limit Exceeded
const JudgeMLE = 8 //Memory Limit Exceeded
const JudgeOLE = 9 //Output Limit Exceeded

const LanguageNA = 0   //None
const LanguageC = 1    //C
const LanguageCPP = 2  //C++
const LanguageJAVA = 3 //Java

const ModuleNA = 0 //None
const ModuleP = 1  //None
const ModuleC = 2  //Contest
const ModuleE = 3  //Exercise

const PrivilegeNA = 0 //None
const PrivilegePU = 1 //Primary User
const PrivilegeSB = 2 //Source Broswer
const PrivilegeAD = 3 //Admin

const EncryptNA = 0 //None
const EncryptPB = 1 //Public
const EncryptPT = 2 //Private
const EncryptPW = 3 //Password
