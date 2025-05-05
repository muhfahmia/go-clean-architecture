package enum

type UserStatus string
type UserRole string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusInactive  UserStatus = "inactive"
	UserStatusSuspended UserStatus = "suspended"
)

const (
	UserRoleSuperadmin       UserRole = "superadmin"
	UserRoleAdmin            UserRole = "admin"
	UserRoleCashier          UserRole = "cashier"
	UserRoleCustomer         UserRole = "customer"
	UserRoleGuest            UserRole = "guest"
	UserRoleManagerInventory UserRole = "manager-inventory"
	UserRoleManagerOrder     UserRole = "manager-order"
	UserRoleManagerUser      UserRole = "manager-user"
	UserRoleManagerReport    UserRole = "manager-report"
	UserRoleManagerFinance   UserRole = "manager-finance"
)
