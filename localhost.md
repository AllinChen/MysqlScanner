
# DATABASE - dbinjection 
### 库 dbinjection 中表信息
| 序号 | 名称 | 来源 | 类型 | 存储引擎 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | db_injection_backup | dbinjection | BASE TABLE | InnoDB | 备份信息表 |
| 2 | db_injection_cluster | dbinjection | BASE TABLE | InnoDB | 集群信息表 |
| 3 | db_injection_exec_item | dbinjection | BASE TABLE | InnoDB | 子任务里的一条sql，执行单位 |
| 4 | db_injection_file | dbinjection | BASE TABLE | InnoDB | 文件信息表 |
| 5 | db_injection_rule | dbinjection | BASE TABLE | InnoDB | 规则信息表 |
| 6 | db_injection_rule_status | dbinjection | BASE TABLE | InnoDB | 规则状态任务表 |
| 7 | db_injection_subtask | dbinjection | BASE TABLE | InnoDB | 子任务表 |
| 8 | db_injection_task | dbinjection | BASE TABLE | InnoDB | 任务表 |
| 9 | projectfile | dbinjection | BASE TABLE | InnoDB | 项目附件表 |
| 10 | view2 | dbinjection | VIEW | null | VIEW |



## TABLE - db_injection_backup 
 表 db_injection_backup 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | data | null | NO | mediumtext | 16777215 | null |  | 备份数据 |
| 3 | rollback_user |  | NO | varchar | 60 | null |  | 恢复执行人 |
| 4 | is_rolled_back | 0 | NO | tinyint | null | 3 |  | 是否已恢复 |
| 5 | ct | 0 | NO | bigint | null | 19 |  | 创建时间 |
| 6 | rollback_time | 0 | NO | bigint | null | 19 |  | 回滚时间 |



## TABLE - db_injection_cluster 
 表 db_injection_cluster 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | name |  | NO | varchar | 160 | null |  | 名称 |
| 3 | description |  | NO | varchar | 200 | null |  | 描述 |
| 4 | addr |  | NO | varchar | 80 | null |  | 访问地址 |
| 5 | user |  | NO | varchar | 120 | null |  | 用户 |
| 6 | pwd |  | NO | varchar | 120 | null |  | 加密后秘钥 |
| 7 | ct | 0 | NO | bigint | null | 19 |  | 创建时间 |
| 8 | ut | 0 | NO | bigint | null | 19 |  | 更改时间 |
| 9 | operator |  | NO | varchar | 40 | null |  | 操作人 |



## TABLE - db_injection_exec_item 
 表 db_injection_exec_item 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | task_id | null | NO | bigint | null | 20 |  | task id |
| 3 | subtask_id | null | NO | bigint | null | 20 |  | subtask id |
| 4 | sql_content | null | YES | text | 65535 | null |  | sql |
| 5 | remark |  | NO | varchar | 800 | null |  | 备注 |
| 6 | affect_rows | 0 | NO | int | null | 10 |  | 影响行数 |
| 7 | rule_comments |  | NO | varchar | 80 | null |  | 规则审核结果 |
| 8 | status |  | NO | varchar | 40 | null |  | 子项状态 |
| 9 | exec_info |  | NO | varchar | 200 | null |  | 执行信息 |
| 10 | backup_status |  | NO | varchar | 40 | null |  | 子项备份状态 |
| 11 | backup_info |  | NO | varchar | 200 | null |  | 备份异常信息 |
| 12 | backup_id | null | NO | bigint | null | 20 |  | backup id |
| 13 | ut | 0 | NO | bigint | null | 19 |  | 更改时间 |
| 14 | et | 0 | NO | bigint | null | 19 |  | 执行时间 |



## TABLE - db_injection_file 
 表 db_injection_file 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | name |  | NO | varchar | 80 | null |  | 文件名称 |
| 3 | location |  | NO | varchar | 800 | null |  | 文件位置 |
| 4 | ct | 0 | NO | bigint | null | 19 |  | 创建时间 |



## TABLE - db_injection_rule 
 表 db_injection_rule 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | item |  | NO | varchar | 80 | null |  | 规则代号 |
| 3 | name | null | YES | varchar | 800 | null |  | 规则名称 |
| 4 | summary |  | NO | varchar | 800 | null |  | 规则摘要 |
| 5 | content | null | YES | varchar | 800 | null |  | 规则解释 |
| 6 | close | 0 | NO | int | null | 10 |  | 规则开关 |
| 7 | level | 0 | NO | int | null | 10 |  | 规则告警等级 |
| 8 | case |  | YES | varchar | 800 | null |  | 规则样例 |
| 9 | position | null | YES | int | null | 10 |  | 字符位置 |
| 10 | checkfunc |  | NO | varchar | 200 | null |  | 规则函数名 |
| 11 | threshold | null | YES | varchar | 200 | null |  | 阈值 |



## TABLE - db_injection_rule_status 
 表 db_injection_rule_status 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | name |  | NO | varchar | 80 | null |  | 规则名 |
| 3 | close | 0 | NO | tinyint | null | 3 |  | 开关 |



## TABLE - db_injection_subtask 
 表 db_injection_subtask 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | task_id | null | NO | bigint | null | 20 |  | task id |
| 3 | task_type |  | NO | varchar | 10 | null |  | 任务类型 |
| 4 | db_name |  | NO | varchar | 120 | null |  | 数据库名称 |
| 5 | cluster_name |  | NO | varchar | 120 | null |  | 数据库集群 |



## TABLE - db_injection_task 
 表 db_injection_task 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | bigint | null | 20 | auto_increment | 自增主键 |
| 2 | name |  | NO | varchar | 40 | null |  | 名称 |
| 3 | status |  | NO | varchar | 40 | null |  | 状态 |
| 4 | creator |  | NO | varchar | 40 | null |  | 创建者 |
| 5 | reviewer |  | NO | varchar | 40 | null |  | 审查者 |
| 6 | executor |  | NO | varchar | 40 | null |  | 执行者 |
| 7 | exec_info |  | NO | varchar | 40 | null |  | 执行信息 |
| 8 | reject_content |  | NO | varchar | 500 | null |  | 驳回信息 |
| 9 | ct | 0 | NO | bigint | null | 19 |  | 创建时间 |
| 10 | ut | 0 | NO | bigint | null | 19 |  | 更改时间 |
| 11 | et | 0 | NO | bigint | null | 19 |  | 执行时间 |
| 12 | ft | 0 | NO | bigint | null | 19 |  | 执行结束时间 |



## TABLE - projectfile 
 表 projectfile 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | null | NO | int | null | 10 | auto_increment | 附件id |
| 2 | fileuploadercode | null | YES | varchar | 128 | null |  | 附件上传者code |
| 3 | projectid | null | YES | int | null | 10 |  | 项目id;此列受project表中的id列约束 |
| 4 | filename | null | YES | varchar | 512 | null |  | 附件名 |
| 5 | fileurl | null | YES | varchar | 512 | null |  | 附件下载地址 |
| 6 | filesize | null | YES | bigint | null | 19 |  | 附件大小，单位Byte |



## TABLE - view2 
 表 view2 中列信息
| 序号 | 名称 | 默认值 | 允许空值 | 数据类型 | 最大长度 | 小数位 | 其他 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 0 | NO | int | null | 10 |  | 附件id |
| 2 | fileuploadercode | null | YES | varchar | 128 | null |  | 附件上传者code |
| 3 | projectid | null | YES | int | null | 10 |  | 项目id;此列受project表中的id列约束 |
| 4 | filename | null | YES | varchar | 512 | null |  | 附件名 |
| 5 | fileurl | null | YES | varchar | 512 | null |  | 附件下载地址 |
| 6 | filesize | null | YES | bigint | null | 19 |  | 附件大小，单位Byte |


