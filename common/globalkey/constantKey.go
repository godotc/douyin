package globalkey

// 软删除只是在数据库表中添加量该字段

var DelStateNo int64 = 0  //未删除
var DelStateYes int64 = 1 //已删除
