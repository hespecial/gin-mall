package e

type Code int

const (
	Success       Code = iota // 响应成功
	InvalidParams             // 参数错误
	UnknownError              // 未知错误

	ErrorUserExists        // 用户已存在
	ErrorEncryptPassword   // 密码加密错误
	ErrorEncryptMoney      // 金额加密错误
	ErrorCreateUser        // 创建用户错误
	ErrorAccountInvalid    // 用户名或密码错误
	ErrorGetUserByID       // 根据id获取用户失败
	ErrorUpdateUser        // 更新用户失败
	ErrorIncorrectPassword // 密码错误
	ErrorUploadAvatar      // 头像上传错误
	ErrorFollowUser        // 关注用户失败
	ErrorUnfollowUser      // 取消关注失败
	ErrorGetFollowingList  // 获取关注列表失败
	ErrorGetFollowerList   // 获取粉丝列表失败

	ErrorGenerateToken // token生成错误
	ErrorParseToken    // token解析错误
	ErrorContextValue  // 上下文值传递错误

	ErrorUploadFile     // 文件上传错误
	ErrorFileError      // 文件错误
	ErrorOSSUploadError // OSS文件上传错误
	ErrorFileType       // 文件类型错误

	ErrorSendEmail            // 发送邮件错误
	ErrorUpdateEmail          // 更新邮箱错误
	ErrorEmailLinkExpire      // 邮件确认链接已过期
	ErrorSendEmailTooFrequent // 邮件发送操作频繁

	ErrorGetCategoryList // 获取商品分类失败
	ErrorGetProductList  // 获取商品列表失败
	ErrorGetProductByID  // 根据ID获取商品失败
	ErrorInvalidIDParam  // 非法的id参数
	ErrorSearchProduct   // 搜索商品失败

	ErrorGetFavoriteList // 获取收藏列表失败
	ErrorAddFavorite     // 收藏商品失败
	ErrorDeleteFavorite  // 取消收藏失败

	ErrorGetCart                // 获取购物车失败
	ErrorAddCartItem            // 添加购物项失败
	ErrorDeleteCartItem         // 删除购物项失败
	ErrorClearCart              // 清空购物车失败
	ErrorUpdateCartItemQuantity // 更新购物项数量失败
	ErrorCacheCartItems         // 缓存购物项失败
	ErrorDeleteCartCache        // 删除购物车缓存失败

	ErrorGetAddressList // 获取用户地址列表失败
	ErrorGetAddressInfo // 获取地址信息失败
	ErrorAddAddress     // 添加用户地址失败
	ErrorUpdateAddress  // 更新用户地址失败
	ErrorDeleteAddress  // 删除用户地址失败

	ErrorGetOrderList         // 获取订单列表失败
	ErrorGetOrderInfo         // 获取订单详情失败
	ErrorCreateOrder          // 创建订单失败
	ErrorCalculateTotalAmount // 计算总金额失败
	ErrorDeleteOrder          // 删除订单失败
)
