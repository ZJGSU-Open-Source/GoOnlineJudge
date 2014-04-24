package config

var ProblemPerPage = 100
var ContestPerPage = 100
var ExercisePerPage = 100
var SolutionPerPage = 50
var UserPerPage = 50

var JudgeNA = 0  //None
var JudgePD = 1  //Pending
var JudgeRJ = 2  //Running & judging
var JudgeAC = 3  //Accepted
var JudgeCE = 4  //Compile Error
var JudgeRE = 5  //Runtime Error
var JudgeWA = 6  //Wrong Answer
var JudgeTLE = 7 //Time Limit Exceeded
var JudgeMLE = 8 //Memory Limit Exceeded
var JudgeOLE = 9 //Output Limit Exceeded

var LanguageNA = 0   //None
var LanguageC = 1    //C
var LanguageCPP = 2  //C++
var LanguageJAVA = 3 //Java

var ModuleNA = 0 //None
var ModuleP = 1  //Problem
var ModuleC = 2  //Contest
var ModuleE = 3  //Exercise

var PrivilegeNA = 0 //None
var PrivilegePU = 1 //Primary User
var PrivilegeSB = 2 //Source Broswer
var PrivilegeAD = 3 //Admin
