# BaseUrl

http://43.140.194.248:8080/api

# 得到所有不良反应报告

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20

**接口地址**：`GET` /adverse

**调用例子**：/adverse?page=1&pageSize=5

# 创建一个不良反应报告

**必选参数**：`token`：令牌

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
            "id": 2,// 疫苗id
            "vaccinateDate": null,// 接种日期
            "dose": "1",// 接种剂次
            "route": "",// 接种途径
            "site": "face"// 接种部位
        }
    ]
}
~~~

**接口地址**：`POST` /adverse

**调用例子**：/adverse?token=...

# 删除一个不良反应报告

**必选参数**：`id`：不良反应报告id，`token`：令牌

**接口地址**：`DELETE` /adverse

**调用例子**：/adverse?id=2&token=...

# 查询疫苗数据

说明：调用此接口，可获得所有疫苗数据。如果productName不为空，则会返回所有与productName详细的疫苗数据。

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`productName`：产品名称

**接口地址**：`GET` /vaccine/cfda

**调用例子**：/vaccine/cfda?page=1&pageSize=5&productName=新冠

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

# 提交一条疫苗数据

说明：调用此接口，可以提交一条疫苗数据。携带数据格式与返回的疫苗数据格式一致。

**接口地址**：`POST` /vaccine/cfda

**调用例子**：/vaccine/cfda

# 删除疫苗

**必选参数**：`id`：删除疫苗id

**接口地址**：`DELETE` /vaccine/cfda

**调用例子**：/vaccine/cfda

# 得到疫苗展示数据

1. **可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认5，`limit`：每个疫苗类型中疫苗的最大数量

   **接口地址**：`GET` /vaccine/example

   **调用例子**：/vaccine/example

2. **可选参数**：`tid`：疫苗类型id，`limit`：每个疫苗类型中疫苗的最大数量

   **接口地址**：`GET` /vaccine/example/:tid

   **调用例子**：/vaccine/example/2

# 用户注册

说明：用户名不能重复

**必选参数**：`username`：用户名，`password`：密码

**接口地址**：`POST` /user/register

**调用例子**：/user/register?username=test&password=123456

# 用户登录

**必选参数**：`username`：用户名，`password`：密码

**接口地址**：`POST` /user/login

**调用例子**：/user/login?username=test&password=123456

# 登录状态检查/刷新

说明：token有效期为24小时

**必选参数**：`token`：令牌

**接口地址**：`GET` /user/status

**调用例子**：/user/status?token=...

# 用户注销

**必选参数**：`token`：令牌

**接口地址**：`GET` /user/logout

**调用例子**：/user/logout?token=...

# 获取用户列表

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20

**接口地址**：`GET` /user/list

**调用例子**：/user/list

# 查询OAE词条

1. **可选参数**：`label`：名称

   **接口地址**：`GET` /oae/label

   **调用例子**：/oae/label?label=fever

2. **必选参数**：`IRI`：OAE词条链接

   **接口地址**：`GET` /oae/IRI

   **调用地址**：/oae/IRI?IRI=http://purl.obolibrary.org/obo/OAE_0000043

# 获取OAE父类词条

**必选参数**：`IRI`：词条链接

**接口地址**：`GET` /oae/parent

**调用例子**：/oae/parent?IRI=http://purl.obolibrary.org/obo/OAE_0000043

# 获取问卷列表

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20

**接口地址**：`GET` /questionnaire

**调用例子**：/questionnaire?page=1&pageSize=20

# 获取用户的所有问卷

**路径参数**：`uid`：用户id

**接口地址**：`GET` /questionnaire/user/:uid

**调用例子**：/questionnaire/user/12

# 创建一个问卷

**必选参数**：`token`：令牌

**携带json**：

~~~json
{
    "name": "dadasdasd",//名称
    "description": ""//描述
}
~~~

**接口地址**：`POST` /questionnaire

**调用例子**：/questionnaire?token=...

# 删除问卷

**必选参数**：`token`：令牌

**路径参数**：`id`：问卷id

**接口地址**：`DELETE` /questionnaire/:id

**调用例子**：/quesionnaire/15?token=...

# 获取问卷中的所有问题

**路径参数**：`id`：问卷id

**接口地址**：`GET`：/questionnaire/:id/questions

**调用例子**：/questionnaire/15/questions

# 获取所有问题类型

**接口地址**：`GET` /questionnaire/questions/type

**调用例子**：/questionnaire/questions/type

# 创建问题

**必选参数**：`token`：令牌

**路径参数**：`id`：问卷id

**携带json**：

~~~json
[{
    "content": "question 1",
    "type": "单选题",
    "isRequired": true,
    "options": [
        "option"
    ],
    "order": 5
},{
    "content": "sadafasfad",
    "type": "填空题",
    "isRequired": false
}]
~~~

**接口地址**：`GET` /questionnaire/:id/questions

**调用例子**：/questionnaire/15/questions?token=...

# 删除问题

**必选参数**：`token`：令牌

**路径参数**：`id`：问卷id，`qid`：问题id

**接口地址**：`DELETE` /questionnaire/:id/questions/:qid

**调用例子**：/questionnaire/15/questions/30?token=...

# 获取一个问卷的所有回答

**必选参数**：`token`：令牌

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认10

**路径参数**：`id`：问卷id

**接口地址**：`GET` /questionnaire/:id/response

**调用例子**：/questionnaire/15/response?token=...&page=1&pageSize=5

# 获取我的所有回答

**必选参数**：`token`：令牌

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认10

**接口地址**：`GET` /questionnaire/response/mine

**调用例子**：/questionnaire/response/mine?token=...&page=1&pageSize=5

# 提交回答

**必选参数**：`token`：令牌

**路径参数**：`id`：问卷id

**携带json**：

~~~json
{
    "responseTermList": [
        {
            "questionId": 29,
            "answer": "test answer"
        }
    ]
}
~~~

**接口地址**：`POST` /questionnaire/:id/response

**调用例子**：/questionnaire/15/response?token=...

# 删除回答

**必选参数**：`token`：令牌，`responseId`：目标回答id

**路径参数**：`id`：问卷id

**接口地址**：`POST` /questionnaire/:id/response

**调用例子**：/questionnaire/15/response?token=...&responseId=16

# 检索Vaers数据

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20

**必选参数**：`vaccineId`：疫苗id，`symptomId`：症状Id，二者必须选填其一

**接口地址**：`GET` /vaers

**调用例子**：/vaers?page=1&pageSize=5&vaccineId=72

# 获取Vaers数据

**路径参数**：`vaersId`：Vaers数据Id

**接口地址**：`GET` /vaers/:vaersId

**调用例子**：/vaers/:2547730

# 获取Vaers的疫苗列表

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`keyword`：疫苗名称

**接口地址**：`GET` /vaers/vaccine

**调用例子**：/vaers/vaccine?page=1&pageSize=5&keyword=COVId

# 获取Vaers的症状列表

**可选参数**：`page`：页码，默认1，`pageSize`：页面大小，默认20，`keyword`：症状名称

**接口地址**：`GET` /vaers/symptom

**调用例子**：/vaers/symptom?page=1&pageSize=5&keyword=EYE
