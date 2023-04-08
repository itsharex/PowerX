// Code generated by goctl. DO NOT EDIT.
package types

type GetEmployeeOptionsRequest struct {
	LikeName        string `form:"likeName,optional"`
	LikeEmail       string `form:"likeEmail,optional"`
	LikePhoneNumber string `form:"likePhoneNumber,optional"`
	PageIndex       int    `form:"pageIndex,optional"`
	PageSize        int    `form:"pageSize,optional"`
}

type EmployeeOption struct {
	Id          int64  `json:"id"`
	Avatar      string `json:"avatar"`
	Account     string `json:"account"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type GetEmployeeOptionsReply struct {
	List      []EmployeeOption `json:"list"`
	PageIndex int              `json:"pageIndex"`
	PageSize  int              `json:"pageSize"`
	Total     int64            `json:"total"`
}

type EmployeeQueryRoleOption struct {
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}

type EmployeeQueryDepartmentOption struct {
	DepartmentId   int64  `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
}

type GetEmployeeQueryOptionsReply struct {
	Positions   []string                        `json:"positions"`
	Roles       []EmployeeQueryRoleOption       `json:"roles"`
	Departments []EmployeeQueryDepartmentOption `json:"departments"`
}

type GetDepartmentOptionsRequest struct {
	Ids       []int64 `form:"ids,optional"`
	LikeName  string  `form:"likeName,optional"`
	PageIndex int     `form:"pageIndex,optional"`
	PageSize  int     `form:"pageSize,optional"`
}

type DepartmentOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GetDepartmentOptionsReply struct {
	List      []DepartmentOption `json:"list"`
	PageIndex int                `json:"pageIndex"`
	PageSize  int                `json:"pageSize"`
	Total     int64              `json:"total"`
}

type DepartmentLeader struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

type DepartmentNode struct {
	Id          int64            `json:"id"`
	DepName     string           `json:"depName"`
	Leader      DepartmentLeader `json:"leader"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Remark      string           `json:"remark"`
	Children    []DepartmentNode `json:"children"`
}

type GetDepartmentTreeRequest struct {
	DepId int64 `path:"depId"`
}

type GetDepartmentTreeReply struct {
	DepTree DepartmentNode `json:"depTree"`
}

type CreateDepartmentRequest struct {
	DepName     string `json:"depName"`
	LeaderId    int64  `json:"leaderId"`
	PId         int64  `json:"pId"`
	Desc        string `json:"desc,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Remark      string `json:"remark,optional"`
}

type CreateDepartmentReply struct {
	Id int64 `json:"id"`
}

type DeleteDepartmentRequest struct {
	Id int64 `path:"id"`
}

type DeleteDepartmentReply struct {
	Id int64 `json:"id"`
}

type PatchDepartmentRequest struct {
	Id          int64  `path:"id"`
	DepName     string `json:"depName,optional"`
	LeaderId    int64  `json:"leaderId,optional"`
	PId         int64  `json:"pId,optional"`
	Desc        string `json:"desc,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Remark      string `json:"remark,optional"`
}

type PatchDepartmentReply struct {
	*Department
}

type Department struct {
	Id          int64            `json:"id"`
	DepName     string           `json:"depName"`
	Leader      DepartmentLeader `json:"leader"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Remark      string           `json:"remark"`
}

type GetDepartmentRequest struct {
	Id int64 `path:"id"`
}

type GetDepartmentReply struct {
	*Department
}

type GetEmployeeRequest struct {
	Id int64 `path:"id"`
}

type GetEmployeeReply struct {
	*Employee
}

type ListEmployeesRequest struct {
	Ids             []int64  `form:"ids,optional"`
	LikeName        string   `form:"likeName,optional"`
	LikeEmail       string   `form:"likeEmail,optional"`
	DepIds          []int64  `form:"depIds,optional"`
	Positions       []string `form:"positions,optional"`
	LikePhoneNumber string   `form:"likePhoneNumber,optional"`
	RoleCodes       []string `form:"roleCodes,optional"`
	IsEnabled       *bool    `form:"isEnable,optional"`
	PageIndex       int      `form:"pageIndex,optional"`
	PageSize        int      `form:"pageSize,optional"`
}

type EmployeeDepartment struct {
	DepId   int64  `json:"depId"`
	DepName string `json:"depName"`
}

type Employee struct {
	Id            int64               `json:"id"`
	Account       string              `json:"account"`
	Name          string              `json:"name"`
	Email         string              `json:"email"`
	MobilePhone   string              `json:"mobilePhone"`
	Gender        string              `json:"gender"`
	NickName      string              `json:"nickName,optional"`
	Desc          string              `json:"desc,optional"`
	Avatar        string              `json:"avatar,optional"`
	ExternalEmail string              `json:"externalEmail,optional"`
	Roles         []string            `json:"roles"`
	Department    *EmployeeDepartment `json:"department"`
	Position      string              `json:"position"`
	JobTitle      string              `json:"jobTitle"`
	IsEnabled     bool                `json:"isEnabled"`
	CreatedAt     string              `json:"createdAt"`
}

type ListEmployeesReply struct {
	List      []Employee `json:"list"`
	PageIndex int        `json:"pageIndex"`
	PageSize  int        `json:"pageSize"`
	Total     int64      `json:"total"`
}

type SyncEmployeesRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type SyncEmployeesReply struct {
	Status bool `json:"status"`
}

type CreateEmployeeRequest struct {
	Account       string `json:"account"`
	Name          string `json:"name"`
	NickName      string `json:"nickName,optional"`
	Desc          string `json:"desc,optional"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar,optional"`
	ExternalEmail string `json:"externalEmail,optional"`
	MobilePhone   string `json:"mobilePhone,optional"`
	Gender        string `json:"gender,options=male|female|un_know"`
	DepId         int64  `json:"depId"`
	Position      string `json:"position,optional"`
	JobTitle      string `json:"jobTitle,optional"`
	Password      string `json:"password,optional"`
}

type CreateEmployeeReply struct {
	Id int64 `json:"id"`
}

type UpdateEmployeeRequest struct {
	Id            int64  `path:"id"`
	Name          string `json:"name,optional"`
	NickName      string `json:"nickName,optional"`
	Desc          string `json:"desc,optional"`
	Email         string `json:"email,optional"`
	Avatar        string `json:"avatar,optional"`
	ExternalEmail string `json:"externalEmail,optional"`
	MobilePhone   string `json:"mobilePhone,optional"`
	Gender        string `json:"gender,optional,options=male|female|un_know"`
	DepId         int64  `json:"depId,optional"`
	Position      string `json:"position,optional"`
	JobTitle      string `json:"jobTitle,optional"`
	Password      string `json:"password,optional"`
	Status        string `json:"status,optional,options=enabled|disabled"`
}

type UpdateEmployeeReply struct {
	*Employee
}

type DeleteEmployeeRequest struct {
	Id int64 `path:"id"`
}

type DeleteEmployeeReply struct {
	Id int64 `json:"id"`
}

type ResetPasswordRequest struct {
	UserId int64 `json:"userId"`
}

type ResetPasswordReply struct {
	Status string `json:"status"`
}

type AdminAPI struct {
	Id        int64  `json:"id"`
	API       string `json:"api"`
	Method    string `json:"method"`
	Name      string `json:"name"`
	GroupId   int64  `json:"groupId"`
	GroupName string `json:"groupName"`
	Desc      string `json:"desc"`
}

type AdminRole struct {
	RoleCode   string     `json:"roleCode"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc"`
	IsReserved bool       `json:"isReserved"`
	APIList    []AdminAPI `json:"apiList"`
	MenuNames  []string   `json:"menuNames"`
}

type ListRolesReply struct {
	List []AdminRole `json:"list"`
}

type CreateRoleRequest struct {
	RoleCode  string   `json:"roleCode"`
	Name      string   `json:"name"`
	Desc      string   `json:"desc"`
	APIIds    []int64  `json:"apiIds"`
	MenuNames []string `json:"menuNames"`
}

type CreateRoleReply struct {
	RoleCode string `json:"roleCode"`
}

type GetRoleRequest struct {
	RoleCode string `path:"roleCode"`
}

type GetRoleReply struct {
	*AdminRole
}

type PatchRoleReqeust struct {
	RoleCode  string   `path:"roleCode"`
	Name      string   `json:"name"`
	Desc      string   `json:"desc,optional"`
	APIIds    []int64  `json:"apiIds,optional"`
	MenuNames []string `json:"menuNames,optional"`
}

type PatchRoleReply struct {
	*AdminRole
}

type SetRolePermissionsRequest struct {
	RoleCode string  `path:"roleCode"`
	APIIds   []int64 `json:"apiIds"`
}

type SetRolePermissionsReply struct {
	Status string `json:"status"`
}

type SetRoleEmployeesRequest struct {
	RoleCode    string  `path:"roleCode"`
	EmployeeIds []int64 `json:"employeeIds"`
}

type SetRoleEmployeesReply struct {
	Status string `json:"status"`
}

type ListAPIRequest struct {
	GroupId int64 `form:"groupId,optional"`
}

type ListAPIReply struct {
	List []AdminAPI `json:"list"`
}

type GetRoleEmployeesReqeust struct {
	RoleCode  string `path:"roleCode"`
	PageIndex int    `form:"pageIndex"`
	PageSize  int    `form:"pageSize"`
}

type RoleEmployeeDepartment struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RoleEmployee struct {
	Id          int64                   `json:"id"`
	Name        string                  `json:"name"`
	Nickname    string                  `json:"nickname"`
	Account     string                  `json:"account"`
	PhoneNumber string                  `json:"phoneNumber"`
	Department  *RoleEmployeeDepartment `json:"department"`
	Email       string                  `json:"email"`
}

type GetRoleEmployeesReply struct {
	List      []RoleEmployee `json:"list"`
	PageIndex int            `json:"pageIndex"`
	PageSize  int            `json:"pageSize"`
	Total     int64          `json:"total"`
}

type SetUserRolesRequest struct {
	UserId    int64    `path:"userId"`
	RoleCodes []string `json:"roleCodes"`
}

type SetUserRolesReply struct {
	Status string `json:"status"`
}

type LoginRequest struct {
	UserName    string `json:"userName,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Password    string `json:"password"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type ExchangeRequest struct {
	Type string `path:"type,optional=wechat"`
	Code string `json:"code"`
}

type ExchangeReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type LeadExternalId struct {
	OpenIDInMiniProgram           string `json:"openIdInMiniProgram,optional"`
	OpenIDInWeChatOfficialAccount string `json:"openIdInWeChatOfficialAccount,optional"`
	OpenIDInWeCom                 string `json:"openIdInWeCom,optional"`
}

type LeadInviter struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

type Lead struct {
	Id        int64       `json:"id"`
	Name      string      `json:"name"`
	Mobile    string      `json:"mobile"`
	Email     string      `json:"email"`
	Inviter   LeadInviter `json:"inviter"`
	Source    string      `json:"source"`
	Status    string      `json:"status"`
	CreatedAt string      `json:"createdAt"`
	LeadExternalId
}

type ListLeadsRequest struct {
	LikeTitle       string   `json:"likeTitle,optional"`
	LikePhoneNumber string   `json:"likePhoneNumber,optional"`
	Sources         []string `json:"sources,optional"`
	Statuses        []string `json:"statuses,optional"`
}

type ListLeadsReply struct {
	List      []Lead `json:"list"`
	PageIndex int32  `json:"pageIndex"`
	PageSize  int32  `json:"pageSize"`
	Total     int64  `json:"total"`
}

type CreateLeadRecord struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile,optional"`
	Email     string `json:"email,optional"`
	InviterId int64  `json:"inviterId,optional"`
	Source    string `json:"source,optional,options=mini_program|tiktok|h5|ad"`
	Status    string `json:"status,optional,options=open|following|closed_lost|closed_won"`
	Type      string `json:"type,optional,options=personal|company"`
	IsActived *bool  `json:"isActived,optional"`
	LeadExternalId
}

type CreateLeadRequest struct {
	List []CreateLeadRecord `json:"list"`
}

type CreateLeadReply struct {
	List      []Lead `json:"list"`
	PageIndex int32  `json:"pageIndex"`
	PageSize  int32  `json:"pageSize"`
	Total     int64  `json:"total"`
}

type PatchLeadRequest struct {
	Id        int64  `path:"id"`
	Name      string `json:"name,optional"`
	Mobile    string `json:"mobile,optional"`
	Email     string `json:"email,optional"`
	InviterId int64  `json:"inviterId,optional"`
	Source    string `json:"source,optional,options=mini_program|tiktok|h5|advertisement"`
	Status    string `json:"status,optional,options=open|following|closed_lost|closed_won"`
	Type      string `json:"type,optional,options=personal|company"`
	IsActived *bool  `json:"isActived,optional"`
	LeadExternalId
}

type PatchLeadReply struct {
	Lead
}

type DeleteLeadRequest struct {
	Id int64 `json:"id"`
}

type DeleteLeadReply struct {
	Id int64 `json:"id"`
}

type AssignLeadToEmployeeRequest struct {
	Id         int64 `path:"id"`
	EmployeeId int64 `json:"employeeId"`
}

type AssignLeadToEmployeeReply struct {
	Lead
}

type CustomerExternalId struct {
	OpenIDInMiniProgram           string `json:"openIdInMiniProgram,optional"`
	OpenIDInWeChatOfficialAccount string `json:"openIdInWeChatOfficialAccount,optional"`
	OpenIDInWeCom                 string `json:"openIdInWeCom,optional"`
}

type CustomerInviter struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

type Customer struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	Mobile      string          `json:"mobile"`
	Email       string          `json:"email"`
	Inviter     CustomerInviter `json:"inviter"`
	Source      string          `json:"source"`
	Type        string          `json:"type"`
	IsActivated bool            `json:"isActivated"`
	CreatedAt   string          `json:"createdAt"`
	CustomerExternalId
}

type GetCustomerReqeuest struct {
	Id string `path:"id"`
}

type GetCustomerReply struct {
	Customer Customer `json:"customer"`
}

type ListCustomersRequest struct {
	LikeName   string   `form:"likeName"`
	Sources    []string `form:"sources"`
	LikeMobile string   `form:"likeMobile"`
	Statuses   []string `form:"statuses"`
	PageIndex  int      `form:"page"`
	PageSize   int      `form:"pageSize"`
}

type ListCustomersReply struct {
	Customers []Customer `json:"customers"`
}

type CreateCustomerRequest struct {
	Name        string `json:"name"`
	Mobile      string `json:"mobile,optional"`
	Email       string `json:"email,optional"`
	InviterId   int64  `json:"inviterId,optional"`
	Source      string `json:"source,optional,options=tiktok|ad|miniprogram|wechat|other"`
	Type        string `json:"type,optional,options=personal|company"`
	IsActivated *bool  `json:"isActivated,optional"`
	CustomerExternalId
}

type CreateCustomerReply struct {
	CustomerId int64 `json:"customerId"`
}

type PatchCustomerRequest struct {
	Id          string `path:"id"`
	Name        string `json:"name,optional"`
	Mobile      string `json:"mobile,optional"`
	Email       string `json:"email,optional"`
	InviterId   int64  `json:"inviterId,optional"`
	Source      string `json:"source,optional,options=tiktok|ad|miniprogram|wechat|other"`
	Type        string `json:"type,optional,options=personal|company"`
	IsActivated *bool  `json:"isActivated,optional"`
	CustomerExternalId
}

type PatchCustomerReply struct {
	Customer Customer `json:"customer"`
}

type DeleteCustomerRequest struct {
	Id string `path:"id"`
}

type DeleteCustomerReply struct {
	CustomerId int64 `json:"customerId"`
}

type AssignCustomerToEmployeeRequest struct {
	Id         string `path:"id"`
	EmployeeId int64  `json:"employeeId"`
}

type AssignCustomerToEmployeeReply struct {
	CustomerId int64 `json:"customerId"`
}

type GetMediaListRequest struct {
	MediaType string   `form:"mediaType,optional"`
	Keys      []string `form:"keys,optional"`
	DescBy    string   `form:"descBy,optional,options=createdAt|updatedAt"`
	PageIndex int      `form:"pageIndex,optional"`
	PageSize  int      `form:"pageSize,optional"`
}

type Media struct {
	Key       string `json:"key"`
	MediaType string `json:"mediaType"`
	Meta      string `json:"meta"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GetMediaListReply struct {
	List      []Media `json:"list"`
	PageIndex int     `json:"pageIndex"`
	PageSize  int     `json:"pageSize"`
	Total     int64   `json:"total"`
}

type CreateMediaUploadRequest struct {
	FileName string `json:"fileName"`
}

type CreateMediaUploadRequestReply struct {
	MediaKey  string `json:"mediaKey"`
	UploadURL string `json:"uploadURL"`
	ExpiresAt int64  `json:"expiresAt"`
}

type CreateOrUpdateMediaRequest struct {
	MediaKey  string `path:"mediaKey"`
	MediaType string `json:"mediaType,optional"`
	Meta      string `json:"meta,optional"`
	Remark    string `json:"remark,optional"`
}

type CreateOrUpdateMediaReply struct {
	MediaKey string `json:"mediaKey"`
}

type GetMediaByKeyRequest struct {
	MediaKey string `path:"mediaKey"`
}

type GetMediaByKeyReply struct {
	*Media
}

type DeleteMediaRequest struct {
	Key string `path:"key"`
}

type DeleteMediaReply struct {
	Key string `json:"key"`
}

type GetDictionaryTypesRequest struct {
	PageIndex int `form:"pageIndex,optional"`
	PageSize  int `form:"pageSize,optional"`
}

type DictionaryType struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetDictionaryTypesReply struct {
	List      []DictionaryType `json:"list"`
	PageIndex int              `json:"pageIndex"`
	PageSize  int              `json:"pageSize"`
	Total     int64            `json:"total"`
}

type CreateDictionaryTypeRequest struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,optional"`
}

type CreateDictionaryTypeReply struct {
	Type string `json:"type"`
}

type UpdateDictionaryTypeRequest struct {
	Type        string `path:"type"`
	Name        string `json:"name,optional"`
	Description string `json:"description,optional"`
}

type UpdateDictionaryTypeReply struct {
	*DictionaryType
}

type DeleteDictionaryTypeRequest struct {
	Type string `path:"type"`
}

type DeleteDictionaryTypeReply struct {
	Type string `json:"type"`
}

type GetDictionaryItemsRequest struct {
	Types     []string `form:"type,optional"`
	PageIndex int      `form:"pageIndex,optional"`
	PageSize  int      `form:"pageSize,optional"`
}

type DictionaryItem struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type GetDictionaryItemsReply struct {
	List      []DictionaryItem `json:"list"`
	PageIndex int              `json:"pageIndex"`
	PageSize  int              `json:"pageSize"`
	Total     int64            `json:"total"`
}

type CreateDictionaryItemRequest struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Sort        int    `json:"sort"`
	Description string `json:"description,optional"`
}

type CreateDictionaryItemReply struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

type UpdateDictionaryItemRequest struct {
	Key         string `path:"key"`
	Type        string `path:"type"`
	Name        string `json:"name,optional"`
	Value       string `json:"value,optional"`
	Sort        int    `json:"sort,optional"`
	Description string `json:"description,optional"`
}

type UpdateDictionaryItemReply struct {
	*DictionaryItem
}

type DeleteDictionaryItemRequest struct {
	Key  string `path:"key"`
	Type string `path:"type"`
}

type DeleteDictionaryItemReply struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

type GetOpportunityListRequest struct {
	Name      string `form:"name,optional"`
	Source    string `form:"source,optional"`
	Type      string `form:"type,optional"`
	Stage     string `form:"stage,optional"`
	DescBy    string `form:"descBy,optional,options=createdAt|updatedAt|closedAt"`
	PageIndex int    `form:"pageIndex,optional"`
	PageSize  int    `form:"pageSize,optional"`
}

type Opportunity struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Requirement string  `json:"requirement"`
	CustomerId  int64   `json:"customerId"`
	Probability float32 `json:"probability"`
	Source      string  `json:"source"`
	Type        string  `json:"type"`
	EmployeeId  int64   `json:"employeeId"`
	Stage       string  `json:"stage"`
	ClosedDate  string  `json:"closedDate"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetOpportunityListReply struct {
	List      []Opportunity `json:"list"`
	PageIndex int           `json:"pageIndex"`
	PageSize  int           `json:"pageSize"`
	Total     int64         `json:"total"`
}

type CreateOpportunityRequest struct {
	Name        string  `json:"name"`
	Requirement string  `json:"requirement"`
	CustomerId  int64   `json:"customerId"`
	Probability float32 `json:"probability,optional"`
	Source      string  `json:"source,options=new_customer|old_customer_new_purchase|old_customer_repurchase|old_customer_upgrade"`
	Type        string  `json:"type,options=trial_requirement|requirement_match|detailed_requirement_analysis|solution_provided|quotation|negotiation|closed_unsuccessful|closed_successful"`
	EmployeeId  int64   `json:"employeeId"`
	Stage       string  `json:"stage"`
}

type CreateOpportunityReply struct {
	Id int64 `json:"id"`
}

type AssignEmployeeToOpportunityRequest struct {
	Id         int64 `path:"id"`
	EmployeeId int64 `json:"employeeId"`
}

type AssignEmployeeToOpportunityReply struct {
	Id int64 `json:"id"`
}

type UpdateOpportunityRequest struct {
	Id          int64   `path:"id"`
	Name        string  `json:"name,optional"`
	Requirement string  `json:"requirement,optional"`
	CustomerId  int64   `json:"customerId,optional"`
	Probability float32 `json:"probability,optional"`
	Source      string  `json:"source,optional,options=new_customer|old_customer_new_purchase|old_customer_repurchase|old_customer_upgrade"`
	Type        string  `json:"type,optional,options=trial_requirement|requirement_match|detailed_requirement_analysis|solution_provided|quotation|negotiation|closed_unsuccessful|closed_successful"`
	EmployeeId  int64   `json:"employeeId,optional"`
	Stage       string  `json:"stage,optional"`
	ClosedDate  string  `json:"closedDate,optional"`
}

type UpdateOpportunityReply struct {
	*Opportunity
}

type DeleteOpportunityRequest struct {
	Id int64 `path:"id"`
}

type DeleteOpportunityReply struct {
	Id int64 `json:"id"`
}

type GetUserInfoReply struct {
	Id            int64    `json:"id"`
	Account       string   `json:"account"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	MobilePhone   string   `json:"mobilePhone"`
	Gender        string   `json:"gender"`
	NickName      string   `json:"nickName"`
	Desc          string   `json:"desc"`
	Avatar        string   `json:"avatar"`
	ExternalEmail string   `json:"externalEmail"`
	DepName       string   `json:"depName"`
	Position      string   `json:"position"`
	JobTitle      string   `json:"jobTitle"`
	IsEnabled     bool     `json:"isEnabled"`
	CreatedAt     string   `json:"createdAt"`
	Roles         []string `json:"roles"`
}

type MenuRoles struct {
	MenuName       string   `json:"menuName"`
	AllowRoleCodes []string `json:"allowRoleCodes"`
}

type GetMenuRolesReply struct {
	MenuRoles []MenuRoles `json:"menuRoles"`
}

type ModifyPasswordReqeust struct {
	Password string `json:"password"`
}

type ContractWayGroupNode struct {
	Id        int64                  `json:"id"`
	GroupName string                 `json:"groupName"`
	Children  []ContractWayGroupNode `json:"children"`
}

type GetContractWayGroupTreeRequest struct {
}

type GetContractWayGroupTreeReply struct {
	GroupTree ContractWayGroupNode `json:"tree"`
}

type ContractWayGroup struct {
	Id        int64  `json:"id"`
	GroupName string `json:"groupName"`
}

type GetContractWayGroupListRequest struct {
	GroupName string `form:"groupName,optional"`
}

type GetContractWayGroupListReply struct {
	Groups []ContractWayGroup `json:"groups"`
}

type GetContractWaysRequest struct {
	EmployeeId int64  `form:"employeeId,optional"`
	Name       string `form:"name,optional"`
	StartDate  string `form:"startDate,optional"`
	EndDate    string `form:"endDate,optional"`
	PageIndex  int    `form:"pageIndex"`
	PageSize   int    `form:"pageSize"`
}

type GetContractWaysReply struct {
	List      []ContractWay `json:"list"`
	PageIndex int           `json:"pageIndex"`
	PageSize  int           `json:"pageSize"`
	Total     int64         `json:"total"`
}

type ContractWay struct {
	Id            int64    `json:"id"`
	Type          int      `json:"type"`
	Scene         int      `json:"scene"`
	Style         string   `json:"style,optional"`
	Remark        string   `json:"remark,optional"`
	SkipVerify    bool     `json:"skipVerify,optional"`
	State         string   `json:"state,optional"`
	Users         []string `json:"users,optional"`
	Parties       []int64  `json:"parties,optional"`
	IsTemp        bool     `json:"isTemp,optional"`
	ExpiresIn     int      `json:"expiresIn,optional"`
	ChatExpiresIn int      `json:"chatExpiresIn,optional"`
	UnionId       string   `json:"unionId,optional"`
	IsExclusive   bool     `json:"isExclusive,optional"`
	Conclusions   string   `json:"conclusions,optional"`
}

type CreateContractWayRequest struct {
	Type          int      `json:"type"`
	Scene         int      `json:"scene"`
	Style         string   `json:"style,optional"`
	Remark        string   `json:"remark,optional"`
	SkipVerify    bool     `json:"skipVerify,optional"`
	State         string   `json:"state,optional"`
	Users         []string `json:"users,optional"`
	Parties       []int64  `json:"parties,optional"`
	IsTemp        bool     `json:"isTemp,optional"`
	ExpiresIn     int      `json:"expiresIn,optional"`
	ChatExpiresIn int      `json:"chatExpiresIn,optional"`
	UnionId       string   `json:"unionId,optional"`
	IsExclusive   bool     `json:"isExclusive,optional"`
	Conclusions   string   `json:"conclusions,optional"`
}

type CreateContractWayReply struct {
	Id int64 `json:"id"`
}

type UpdateContractWayRequest struct {
	Id            int64    `path:"id"`
	Type          int      `json:"type,optional"`
	Scene         int      `json:"scene,optional"`
	Style         string   `json:"style,optional"`
	Remark        string   `json:"remark,optional"`
	SkipVerify    bool     `json:"skipVerify,optional"`
	State         string   `json:"state,optional"`
	Users         []string `json:"users,optional"`
	Parties       []int64  `json:"parties,optional"`
	IsTemp        bool     `json:"isTemp,optional"`
	ExpiresIn     int      `json:"expiresIn,optional"`
	ChatExpiresIn int      `json:"chatExpiresIn,optional"`
	UnionId       string   `json:"unionId,optional"`
	IsExclusive   bool     `json:"isExclusive,optional"`
	Conclusions   string   `json:"conclusions,optional"`
}

type UpdateContractWayReply struct {
	ContractWayUpdated ContractWay `json:"contractWayUpdated"`
}

type DeleteContractWayRequest struct {
	Id int64 `path:"id"`
}

type DeleteContractWayReply struct {
	Id int64 `json:"id"`
}

type WeWorkCustomer struct {
	Name            string   `json:"name"`
	AdderId         int64    `json:"adderId"`
	AddTime         string   `json:"addTime"`
	PatchTime       string   `json:"updateTime"`
	AddChannel      string   `json:"addChannel"`
	TagGroupIdList  []int64  `json:"tagGroupIdList"`
	TagIdList       []int64  `json:"tagIdList"`
	PersonalTagList []string `json:"personalTagList"`
	Age             int      `json:"age"`
	Email           string   `json:"email"`
	PhoneNumber     string   `json:"phoneNumber"`
	Address         string   `json:"address"`
	Birthday        string   `json:"birthday"`
	Remark          string   `json:"remark"`
	GroupChatId     int64    `json:"groupChatId"`
}

type GetWeWorkCustomerRequest struct {
	Id string `path:"id"`
}

type GetWeWorkCustomerReply struct {
	WeWorkCustomer
}

type ListWeWorkCustomersRequest struct {
	LikeName     string `form:"likeName,optional"`
	FollowUserId string `form:"followUserId,optional"`
	AddWay       int    `form:"addWay,optional"`
	PageIndex    int    `form:"pageIndex,optional"`
	PageSize     int    `form:"pageSize,optional"`
}

type ListWeWorkCustomersReply struct {
	List      []GetWeWorkCustomerReply `json:"list"`
	PageIndex int                      `json:"pageIndex"`
	PageSize  int                      `json:"pageSize"`
	Total     int64                    `json:"total"`
}

type PatchWeWorkCustomerRequest struct {
	Id          string `path:"id"`
	Name        string `json:"name,optional"`
	Age         int    `json:"age,optional"`
	Email       string `json:"email,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Address     string `json:"address,optional"`
	Birthday    string `json:"birthday,optional"`
	Remark      string `json:"remark,optional"`
	GroupChatId int64  `json:"groupChatId,optional"`
}

type PatchWeWorkCustomerReply struct {
	WeWorkCustomer
}

type SyncWeWorkCustomerReply struct {
	Status string `json:"status"`
}

type SyncWeWorkContactReply struct {
	Status string `json:"status"`
}

type ListWeWorkEmployeeReqeust struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}

type ListWeWorkEmployeeReply struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

type MPCustomerLoginRequest struct {
	Code string `json:"code"`
}

type MPCustomerAuthRequest struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type MPCustomerLoginAuthReply struct {
	OpenID      string `json:"openID"`
	UnionID     string `json:"unionID"`
	PhoneNumber string `json:"phoneNumber"`
	NickName    string `json:"nickName"`
	AvatarURL   string `json:"avatarURL"`
	Gender      string `json:"gender"`
	Token       Token  `json:"token"`
}

type Token struct {
	TokenType    string `json:"tokenType"`
	ExpiresIn    string `json:"expiresIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
