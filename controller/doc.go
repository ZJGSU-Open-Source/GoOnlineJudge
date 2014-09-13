package controller

// controller包中各个控件命名必须是：XXXController，
// 必须实现class.Router接口，并且其方法Route的接收者必须是值类型控件，
// 例如(this XXController)Route(w http.ResponseWriter, r *http.Request)
// Route方法作为路由和控件初始化之用。
// 其余方法应该在Route方法中调用，或者能通过Route方法间接调用。
