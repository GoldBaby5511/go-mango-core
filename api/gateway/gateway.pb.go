// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: gateway.proto

package gateway

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CMDGateway int32

const (
	CMDGateway_IDPulseReq        CMDGateway = 1 //测速请求
	CMDGateway_IDPulseRsp        CMDGateway = 2 //测速回复
	CMDGateway_IDTransferDataReq CMDGateway = 3 //数据转发请求
	CMDGateway_IDTransferDataRsp CMDGateway = 4 //数据转发回复
	CMDGateway_IDAuthInfo        CMDGateway = 5 //认证信息
	CMDGateway_IDHelloReq        CMDGateway = 6 //握手请求
	CMDGateway_IDHelloRsp        CMDGateway = 7 //握手回复
)

// Enum value maps for CMDGateway.
var (
	CMDGateway_name = map[int32]string{
		1: "IDPulseReq",
		2: "IDPulseRsp",
		3: "IDTransferDataReq",
		4: "IDTransferDataRsp",
		5: "IDAuthInfo",
		6: "IDHelloReq",
		7: "IDHelloRsp",
	}
	CMDGateway_value = map[string]int32{
		"IDPulseReq":        1,
		"IDPulseRsp":        2,
		"IDTransferDataReq": 3,
		"IDTransferDataRsp": 4,
		"IDAuthInfo":        5,
		"IDHelloReq":        6,
		"IDHelloRsp":        7,
	}
)

func (x CMDGateway) Enum() *CMDGateway {
	p := new(CMDGateway)
	*p = x
	return p
}

func (x CMDGateway) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDGateway) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_proto_enumTypes[0].Descriptor()
}

func (CMDGateway) Type() protoreflect.EnumType {
	return &file_gateway_proto_enumTypes[0]
}

func (x CMDGateway) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMDGateway) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMDGateway(num)
	return nil
}

// Deprecated: Use CMDGateway.Descriptor instead.
func (CMDGateway) EnumDescriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

type AuthInfo_OpType int32

const (
	AuthInfo_Bind       AuthInfo_OpType = 0 //userId与connId绑定
	AuthInfo_Disconnect AuthInfo_OpType = 1 //断开当前connId连接
)

// Enum value maps for AuthInfo_OpType.
var (
	AuthInfo_OpType_name = map[int32]string{
		0: "Bind",
		1: "Disconnect",
	}
	AuthInfo_OpType_value = map[string]int32{
		"Bind":       0,
		"Disconnect": 1,
	}
)

func (x AuthInfo_OpType) Enum() *AuthInfo_OpType {
	p := new(AuthInfo_OpType)
	*p = x
	return p
}

func (x AuthInfo_OpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthInfo_OpType) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_proto_enumTypes[1].Descriptor()
}

func (AuthInfo_OpType) Type() protoreflect.EnumType {
	return &file_gateway_proto_enumTypes[1]
}

func (x AuthInfo_OpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AuthInfo_OpType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AuthInfo_OpType(num)
	return nil
}

// Deprecated: Use AuthInfo_OpType.Descriptor instead.
func (AuthInfo_OpType) EnumDescriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{4, 0}
}

type HelloRsp_RspFlag int32

const (
	HelloRsp_UNKNOWN       HelloRsp_RspFlag = 0 //未知
	HelloRsp_EncryptInfo   HelloRsp_RspFlag = 1 //加密信息   encrypt_key 这是存在的
	HelloRsp_AdviceNewGate HelloRsp_RspFlag = 2 //建议去新的gate，这时gate_address 必须有内容
	HelloRsp_LoginToken    HelloRsp_RspFlag = 4 //登录令牌
)

// Enum value maps for HelloRsp_RspFlag.
var (
	HelloRsp_RspFlag_name = map[int32]string{
		0: "UNKNOWN",
		1: "EncryptInfo",
		2: "AdviceNewGate",
		4: "LoginToken",
	}
	HelloRsp_RspFlag_value = map[string]int32{
		"UNKNOWN":       0,
		"EncryptInfo":   1,
		"AdviceNewGate": 2,
		"LoginToken":    4,
	}
)

func (x HelloRsp_RspFlag) Enum() *HelloRsp_RspFlag {
	p := new(HelloRsp_RspFlag)
	*p = x
	return p
}

func (x HelloRsp_RspFlag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HelloRsp_RspFlag) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_proto_enumTypes[2].Descriptor()
}

func (HelloRsp_RspFlag) Type() protoreflect.EnumType {
	return &file_gateway_proto_enumTypes[2]
}

func (x HelloRsp_RspFlag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *HelloRsp_RspFlag) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = HelloRsp_RspFlag(num)
	return nil
}

// Deprecated: Use HelloRsp_RspFlag.Descriptor instead.
func (HelloRsp_RspFlag) EnumDescriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{6, 0}
}

//测速请求
type PulseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MySpeed      *uint32 `protobuf:"varint,1,opt,name=my_speed,json=mySpeed" json:"my_speed,omitempty"`
	SpeedData    *uint32 `protobuf:"varint,2,opt,name=speed_data,json=speedData" json:"speed_data,omitempty"`
	AttachedData []byte  `protobuf:"bytes,3,opt,name=attached_data,json=attachedData" json:"attached_data,omitempty"`
}

func (x *PulseReq) Reset() {
	*x = PulseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulseReq) ProtoMessage() {}

func (x *PulseReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulseReq.ProtoReflect.Descriptor instead.
func (*PulseReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *PulseReq) GetMySpeed() uint32 {
	if x != nil && x.MySpeed != nil {
		return *x.MySpeed
	}
	return 0
}

func (x *PulseReq) GetSpeedData() uint32 {
	if x != nil && x.SpeedData != nil {
		return *x.SpeedData
	}
	return 0
}

func (x *PulseReq) GetAttachedData() []byte {
	if x != nil {
		return x.AttachedData
	}
	return nil
}

//测速请求
type PulseRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpeedData   *uint32 `protobuf:"varint,2,opt,name=speed_data,json=speedData" json:"speed_data,omitempty"`
	AttachdData []byte  `protobuf:"bytes,3,opt,name=attachd_data,json=attachdData" json:"attachd_data,omitempty"`
}

func (x *PulseRsp) Reset() {
	*x = PulseRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulseRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulseRsp) ProtoMessage() {}

func (x *PulseRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulseRsp.ProtoReflect.Descriptor instead.
func (*PulseRsp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *PulseRsp) GetSpeedData() uint32 {
	if x != nil && x.SpeedData != nil {
		return *x.SpeedData
	}
	return 0
}

func (x *PulseRsp) GetAttachdData() []byte {
	if x != nil {
		return x.AttachdData
	}
	return nil
}

//数据转发请求
type TransferDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestApptype  *uint32 `protobuf:"varint,1,opt,name=dest_apptype,json=destApptype" json:"dest_apptype,omitempty"` //目标或源apptype
	DestAppid    *uint32 `protobuf:"varint,2,opt,name=dest_appid,json=destAppid" json:"dest_appid,omitempty"`       //目标或源appid
	DataApptype  *uint32 `protobuf:"varint,3,opt,name=data_apptype,json=dataApptype" json:"data_apptype,omitempty"`
	DataCmdid    *uint32 `protobuf:"varint,4,opt,name=data_cmdid,json=dataCmdid" json:"data_cmdid,omitempty"`
	Data         []byte  `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
	ReqId        *uint32 `protobuf:"varint,6,opt,name=req_id,json=reqId" json:"req_id,omitempty"`
	ClientIpV4   *uint32 `protobuf:"varint,7,opt,name=client_ip_v4,json=clientIpV4" json:"client_ip_v4,omitempty"`     //客户端的ip
	AttSessionid *uint64 `protobuf:"varint,8,opt,name=att_sessionid,json=attSessionid" json:"att_sessionid,omitempty"` //联联的session id ，目前只由gate->client
	Gateconnid   *uint64 `protobuf:"varint,9,opt,name=gateconnid" json:"gateconnid,omitempty"`                         //关联的gate连接id
	Gateid       *uint32 `protobuf:"varint,10,opt,name=gateid" json:"gateid,omitempty"`                                //关联的gate_id
	UserId       *uint64 `protobuf:"varint,11,opt,name=user_id,json=userId" json:"user_id,omitempty"`                  //用户ID
}

func (x *TransferDataReq) Reset() {
	*x = TransferDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferDataReq) ProtoMessage() {}

func (x *TransferDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferDataReq.ProtoReflect.Descriptor instead.
func (*TransferDataReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *TransferDataReq) GetDestApptype() uint32 {
	if x != nil && x.DestApptype != nil {
		return *x.DestApptype
	}
	return 0
}

func (x *TransferDataReq) GetDestAppid() uint32 {
	if x != nil && x.DestAppid != nil {
		return *x.DestAppid
	}
	return 0
}

func (x *TransferDataReq) GetDataApptype() uint32 {
	if x != nil && x.DataApptype != nil {
		return *x.DataApptype
	}
	return 0
}

func (x *TransferDataReq) GetDataCmdid() uint32 {
	if x != nil && x.DataCmdid != nil {
		return *x.DataCmdid
	}
	return 0
}

func (x *TransferDataReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TransferDataReq) GetReqId() uint32 {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return 0
}

func (x *TransferDataReq) GetClientIpV4() uint32 {
	if x != nil && x.ClientIpV4 != nil {
		return *x.ClientIpV4
	}
	return 0
}

func (x *TransferDataReq) GetAttSessionid() uint64 {
	if x != nil && x.AttSessionid != nil {
		return *x.AttSessionid
	}
	return 0
}

func (x *TransferDataReq) GetGateconnid() uint64 {
	if x != nil && x.Gateconnid != nil {
		return *x.Gateconnid
	}
	return 0
}

func (x *TransferDataReq) GetGateid() uint32 {
	if x != nil && x.Gateid != nil {
		return *x.Gateid
	}
	return 0
}

func (x *TransferDataReq) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

//数据转发请求
type TransferDataRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *uint32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	ReqId  *uint32 `protobuf:"varint,6,opt,name=req_id,json=reqId" json:"req_id,omitempty"`
}

func (x *TransferDataRsp) Reset() {
	*x = TransferDataRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferDataRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferDataRsp) ProtoMessage() {}

func (x *TransferDataRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferDataRsp.ProtoReflect.Descriptor instead.
func (*TransferDataRsp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *TransferDataRsp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *TransferDataRsp) GetReqId() uint32 {
	if x != nil && x.ReqId != nil {
		return *x.ReqId
	}
	return 0
}

//连接认证信息
type AuthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     *uint64          `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`                                 //用户id
	Gateconnid *uint64          `protobuf:"varint,2,opt,name=gateconnid" json:"gateconnid,omitempty"`                                       //关联的gate连接id
	Result     *uint32          `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`                                               //结果 0成功
	Info       *string          `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`                                                    //描述信息
	OpType     *AuthInfo_OpType `protobuf:"varint,5,opt,name=op_type,json=opType,enum=bs.gateway.AuthInfo_OpType" json:"op_type,omitempty"` //操作类型
}

func (x *AuthInfo) Reset() {
	*x = AuthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfo) ProtoMessage() {}

func (x *AuthInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInfo.ProtoReflect.Descriptor instead.
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *AuthInfo) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *AuthInfo) GetGateconnid() uint64 {
	if x != nil && x.Gateconnid != nil {
		return *x.Gateconnid
	}
	return 0
}

func (x *AuthInfo) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *AuthInfo) GetInfo() string {
	if x != nil && x.Info != nil {
		return *x.Info
	}
	return ""
}

func (x *AuthInfo) GetOpType() AuthInfo_OpType {
	if x != nil && x.OpType != nil {
		return *x.OpType
	}
	return AuthInfo_Bind
}

//握手请求
type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId          *uint32 `protobuf:"varint,2,opt,name=ad_id,json=adId" json:"ad_id,omitempty"`
	Others        *string `protobuf:"bytes,3,opt,name=others" json:"others,omitempty"`
	BuilderNo     *uint32 `protobuf:"varint,4,opt,name=builder_no,json=builderNo" json:"builder_no,omitempty"`
	GameKind      *uint32 `protobuf:"varint,5,opt,name=game_kind,json=gameKind" json:"game_kind,omitempty"`
	ClientVersion *string `protobuf:"bytes,6,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
	ClientType    *uint32 `protobuf:"varint,7,opt,name=client_type,json=clientType" json:"client_type,omitempty"`
	PublicKey     *string `protobuf:"bytes,8,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	EncryptKey    *string `protobuf:"bytes,9,opt,name=encrypt_key,json=encryptKey" json:"encrypt_key,omitempty"` //加密key
	Guid          *string `protobuf:"bytes,10,opt,name=guid" json:"guid,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *HelloReq) GetAdId() uint32 {
	if x != nil && x.AdId != nil {
		return *x.AdId
	}
	return 0
}

func (x *HelloReq) GetOthers() string {
	if x != nil && x.Others != nil {
		return *x.Others
	}
	return ""
}

func (x *HelloReq) GetBuilderNo() uint32 {
	if x != nil && x.BuilderNo != nil {
		return *x.BuilderNo
	}
	return 0
}

func (x *HelloReq) GetGameKind() uint32 {
	if x != nil && x.GameKind != nil {
		return *x.GameKind
	}
	return 0
}

func (x *HelloReq) GetClientVersion() string {
	if x != nil && x.ClientVersion != nil {
		return *x.ClientVersion
	}
	return ""
}

func (x *HelloReq) GetClientType() uint32 {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return 0
}

func (x *HelloReq) GetPublicKey() string {
	if x != nil && x.PublicKey != nil {
		return *x.PublicKey
	}
	return ""
}

func (x *HelloReq) GetEncryptKey() string {
	if x != nil && x.EncryptKey != nil {
		return *x.EncryptKey
	}
	return ""
}

func (x *HelloReq) GetGuid() string {
	if x != nil && x.Guid != nil {
		return *x.Guid
	}
	return ""
}

//握手回复
type HelloRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RspFlag     *uint32  `protobuf:"varint,1,opt,name=rsp_flag,json=rspFlag" json:"rsp_flag,omitempty"`            //通知的消息内容
	GateAddress []string `protobuf:"bytes,2,rep,name=gate_address,json=gateAddress" json:"gate_address,omitempty"` //当前最新的gate地址
	LoginToken  *uint32  `protobuf:"varint,4,opt,name=login_token,json=loginToken" json:"login_token,omitempty"`   //登录令牌
	PublicKey   *string  `protobuf:"bytes,8,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	EncryptKey  *string  `protobuf:"bytes,3,opt,name=encrypt_key,json=encryptKey" json:"encrypt_key,omitempty"` //加密key
	//
	//=0 表示是最新版本
	//=1 表示有新版本,但当前版本还可以用
	//=2 表示老版本必须更新了,当前连接会被断开的
	VersionResult *uint32 `protobuf:"varint,5,opt,name=version_result,json=versionResult" json:"version_result,omitempty"`
	//
	//如果有新的版本,下载地址(一般用于手机)
	DownUrl *string `protobuf:"bytes,6,opt,name=down_url,json=downUrl" json:"down_url,omitempty"`
	//如果Req的guid为空，则这里为其创建一个guid
	Guid *string `protobuf:"bytes,7,opt,name=guid" json:"guid,omitempty"`
}

func (x *HelloRsp) Reset() {
	*x = HelloRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRsp) ProtoMessage() {}

func (x *HelloRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRsp.ProtoReflect.Descriptor instead.
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *HelloRsp) GetRspFlag() uint32 {
	if x != nil && x.RspFlag != nil {
		return *x.RspFlag
	}
	return 0
}

func (x *HelloRsp) GetGateAddress() []string {
	if x != nil {
		return x.GateAddress
	}
	return nil
}

func (x *HelloRsp) GetLoginToken() uint32 {
	if x != nil && x.LoginToken != nil {
		return *x.LoginToken
	}
	return 0
}

func (x *HelloRsp) GetPublicKey() string {
	if x != nil && x.PublicKey != nil {
		return *x.PublicKey
	}
	return ""
}

func (x *HelloRsp) GetEncryptKey() string {
	if x != nil && x.EncryptKey != nil {
		return *x.EncryptKey
	}
	return ""
}

func (x *HelloRsp) GetVersionResult() uint32 {
	if x != nil && x.VersionResult != nil {
		return *x.VersionResult
	}
	return 0
}

func (x *HelloRsp) GetDownUrl() string {
	if x != nil && x.DownUrl != nil {
		return *x.DownUrl
	}
	return ""
}

func (x *HelloRsp) GetGuid() string {
	if x != nil && x.Guid != nil {
		return *x.Guid
	}
	return ""
}

var File_gateway_proto protoreflect.FileDescriptor

var file_gateway_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x62, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x69, 0x0a, 0x08, 0x50,
	0x75, 0x6c, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x79, 0x5f, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x79, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x70, 0x65, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x08, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x70, 0x65, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xd8, 0x02, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x70, 0x70, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x64, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x70, 0x70, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x41, 0x70, 0x70, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6d, 0x64, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6d, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x70, 0x5f, 0x76, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x56, 0x34, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x61, 0x74, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x61, 0x74, 0x65, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x67, 0x61, 0x74, 0x65, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x40, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65,
	0x71, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49,
	0x64, 0x22, 0xc9, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x63,
	0x6f, 0x6e, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x67, 0x61, 0x74,
	0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x22, 0x0a, 0x06, 0x4f, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x69, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x01, 0x22, 0x8f, 0x02,
	0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x67,
	0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x22,
	0xcb, 0x02, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x73, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x73, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x67,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x77, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69,
	0x64, 0x22, 0x4a, 0x0a, 0x07, 0x52, 0x73, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x64,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x04, 0x2a, 0x8a, 0x01,
	0x0a, 0x0a, 0x43, 0x4d, 0x44, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x44, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x52, 0x65, 0x71, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x44, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x52, 0x73, 0x70, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x49, 0x44, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x44, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44,
	0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x10, 0x07, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
}

var (
	file_gateway_proto_rawDescOnce sync.Once
	file_gateway_proto_rawDescData = file_gateway_proto_rawDesc
)

func file_gateway_proto_rawDescGZIP() []byte {
	file_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_proto_rawDescData)
	})
	return file_gateway_proto_rawDescData
}

var file_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gateway_proto_goTypes = []interface{}{
	(CMDGateway)(0),         // 0: bs.gateway.CMDGateway
	(AuthInfo_OpType)(0),    // 1: bs.gateway.AuthInfo.OpType
	(HelloRsp_RspFlag)(0),   // 2: bs.gateway.HelloRsp.RspFlag
	(*PulseReq)(nil),        // 3: bs.gateway.PulseReq
	(*PulseRsp)(nil),        // 4: bs.gateway.PulseRsp
	(*TransferDataReq)(nil), // 5: bs.gateway.TransferDataReq
	(*TransferDataRsp)(nil), // 6: bs.gateway.TransferDataRsp
	(*AuthInfo)(nil),        // 7: bs.gateway.AuthInfo
	(*HelloReq)(nil),        // 8: bs.gateway.HelloReq
	(*HelloRsp)(nil),        // 9: bs.gateway.HelloRsp
}
var file_gateway_proto_depIdxs = []int32{
	1, // 0: bs.gateway.AuthInfo.op_type:type_name -> bs.gateway.AuthInfo.OpType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gateway_proto_init() }
func file_gateway_proto_init() {
	if File_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulseReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulseRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferDataReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferDataRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_proto_depIdxs,
		EnumInfos:         file_gateway_proto_enumTypes,
		MessageInfos:      file_gateway_proto_msgTypes,
	}.Build()
	File_gateway_proto = out.File
	file_gateway_proto_rawDesc = nil
	file_gateway_proto_goTypes = nil
	file_gateway_proto_depIdxs = nil
}
