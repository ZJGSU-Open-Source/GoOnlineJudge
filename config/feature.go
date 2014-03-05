package config

var ModuleProblem = true
var ModuleContest = true
var ModuleExercise = true

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

var SpecialNA = 0 //None
var SpecialST = 1 //Standard
var SpecialSP = 2 //Special
