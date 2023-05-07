# BaseUrl

http://43.140.194.248/api

# 得到不良反应报告

说明：调用此接口，可以获取不良反应报告，可以通过`id`与`uid`获取相应的不良反应报告。

**可选参数**： `page`：页码，默认1，`pageSize`：页面大小，默认20，`id`：指定报告id，`uid`：指定用户

**接口地址**：`GET` /adverse

**调用例子**：:pushpin:/adverse?page=1&pageSize=5	:pushpin:/adverse?uid=23	:pushpin:/adverse?id=22



# 创建一个不良反应报告

说明：调用此接口，可以创建一个不良反应报告记录，可以选择是否携带token。报告中只有不良反应描述是必填的。

**可选参数**：`token`：令牌

**携带json**：

~~~json
{
    "code": "987654321",// 编码
    "name": "jack",// 姓名
    "sex": "N",// 性别，F（男）M（女）N（未知）
    "birth": null,// 出生日期
    "phone": "12345678910",// 联系电话
    "address": "unnnnnnn",// 现住址
    "onsetDate": "1970-01-20T10:46:04Z",// 反应发生日期
    "description": "dasdasdasdad",// 不良反应描述，必填
    "treatmentDepartment": "hhh",// 就诊单位
    "rapporteur": "jjj",// 报告人
    "rapporteurPhone": "kkk",// 报告人联系电话
    "rapporteurAddress": "llll",// 报告单位
    "vaccineList": [// 疫苗接种情况
        {
            "vaccineId": 2,// 疫苗id
            "vaccinateDate": null,// 接种日期
            "dose": "1",// 接种剂次
            "route": "",// 接种途径
            "site": "face"// 接种部位
        }
    ],
    "symptomList": [//不良反应症状列表
        {
            "symptom": "fever",
            "oaeId": 4
        }
    ]
}
~~~

**接口地址**：`POST` /adverse

**调用例子**：:pushpin:/adverse?token=...	:pushpin:/adverse

# 删除一个不良反应报告

**必选参数**：`id`：不良反应报告id，`token`：令牌

**接口地址**：`DELETE` /adverse

**调用例子**：:pushpin:/adverse?id=2&token=...

# 查询疫苗数据

说明：调用此接口，可获得所有疫苗数据。如果productName不为空，则会返回所有与productName详细的疫苗数据。

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`productName`：产品名称，`type`：疫苗类型

**接口地址**：`GET` /vaccine/cfda

**调用例子**：:pushpin:/vaccine/cfda?page=1&pageSize=5&productName=新冠	:pushpin:/vaccine/cfda?type=新冠疫苗

**字段说明**：

~~~json
"id": 1,
"type": "流感疫苗",//疫苗类型
"registerNumber": "国药准字S20080005",//批准文号
"productName": "大流行流感病毒灭活疫苗",//产品名称
"englishName": "Pandemic Influenza Vaccine(Inactivated adjuvanted)",//英文名称
"tradeName": "盼尔来福（Panflu)",//商品名
"dosage": "注射剂",//剂型
"specification": "0.5ml。每一次人用剂量为0.5ml,含大流行流感病毒抗原10μg。",//规格
"owner": "",//上市许可持有人
"ownerAddress": "",//上市许可持有人地址
"productionCompany": "北京科兴生物制品有限公司",//生产单位
"approvalDate": "2018/06/28",//批准日期
"productionAddress": "北京市海淀区上地西路39号，北京市昌平区中关村科技园区昌平园智通路15号",//生产地址
"productionClass": "生物制品",//产品类型
"originalNumber": "",//原批准文号
"drugCode": "86900080000078",//药品本位码
"drugCodeNote": ""//药品本位码备注
~~~

# 获取疫苗类型

说明：如果有id参数，且id参数合法，将会返回目标疫苗类型的详细信息。

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`id`：类型id

**接口地址**：`GET` /vaccine/type

**调用例子**：:pushpin:/vaccine/type?page=1&pageSize=5

# 用户注册

说明：用户名不能重复，可以选择json返回，或者返回cookie。

**可选参数**：`type`：返回方式，可选json/cookie，默认json

**必选参数**：`username`：用户名，`password`：密码

**接口地址**：`POST` /user/register

**调用例子**：:pushpin:/user/register?username=test&password=123456	:pushpin:/user/register?username=test& password=123456&type=cookie

# 用户登录

说明：如果已经登录将会返回一个新的token，token有效期为24小时。服务端首先会检查，是否已经携带了token，如果有会优先使用token进行登录，优先级为cookie>token>password。注意：如果使用token或者password登录，则默认返回json，如果使用cookie登录，则默认返回cookie。

**可选参数**：`username`：用户名，`password`：密码，`token`：已经登录的令牌，`type`：返回方式，可选json/cookie

**接口地址**：`POST` /user/login

**调用例子**：:pushpin:/user/login?username=test&password=123456 :pushpin:/user/login?token=...	:pushpin:/user/login	:pushpin:/user/login?type=json

# 用户注销

说明：注销用户将会删除用户数据，可以通过cookie或者token注销。

**可选参数**：`token`：令牌

**接口地址**：`GET` /user/logout

**调用例子**：:pushpin:/user/logout?token=...	:pushpin:/user/logout

# 查询用户

说明：调用此接口将会通过username查询相似用户。

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`username`：用户名

**接口地址**：`GET` /user/list

**调用例子**：:pushpin:/user/search	:pushpin:/user/search?username=test

# 查询OAE词条

1. **可选参数**：`label`：名称

   **接口地址**：`GET` /oae/label

   **调用例子**：:pushpin:/oae/label?label=fever

2. **必选参数**：`IRI`：OAE词条链接

   **接口地址**：`GET` /oae/IRI

   **调用地址**：:pushpin:/oae/IRI?IRI=http://purl.obolibrary.org/obo/OAE_0000043

# 获取OAE父类词条

**必选参数**：`IRI`：词条链接

**接口地址**：`GET` /oae/parent

**调用例子**：:pushpin:/oae/parent?IRI=http://purl.obolibrary.org/obo/OAE_0000043

# 检索Vaers统计数据

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20

**必选参数**：`vaccineId`：疫苗id，`symptomId`：症状Id，二者必须选填其一

**接口地址**：`GET` /vaers

**调用例子**：:pushpin:/vaers?page=1&pageSize=5&vaccineId=72

# 获取Vaers的疫苗列表

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`keyword`：疫苗名称

**接口地址**：`GET` /vaers/vaccine

**调用例子**：:pushpin:/vaers/vaccine?page=1&pageSize=5&keyword=COVId

# 获取Vaers的症状列表

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`keyword`：症状名称

**接口地址**：`GET` /vaers/symptom

**调用例子**：:pushpin:/vaers/symptom?page=1&pageSize=5&keyword=EYE
