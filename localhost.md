# DATABASE - dbinjection 
### 库 dbinjection 中表信息
| 序号 | 名称 | 来源 | 类型 | 存储引擎 | 创建时间 | 修改时间 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | db_injection_backup | dbinjection | BASE TABLE | InnoDB | null | null | 备份信息表 |
| 2 | db_injection_cluster | dbinjection | BASE TABLE | InnoDB | null | null | 集群信息表 |
| 3 | db_injection_exec_item | dbinjection | BASE TABLE | InnoDB | null | null | 子任务里的一条sql，执行单位 |
| 4 | db_injection_file | dbinjection | BASE TABLE | InnoDB | null | null | 文件信息表 |
| 5 | db_injection_rule | dbinjection | BASE TABLE | InnoDB | null | null | 规则信息表 |
| 6 | db_injection_rule_status | dbinjection | BASE TABLE | InnoDB | null | null | 规则状态任务表 |
| 7 | db_injection_subtask | dbinjection | BASE TABLE | InnoDB | null | null | 子任务表 |
| 8 | db_injection_task | dbinjection | BASE TABLE | InnoDB | null | null | 任务表 |
| 9 | projectfile | dbinjection | BASE TABLE | InnoDB | null | null | 项目附件表 |
| 10 | view2 | dbinjection | VIEW | null | null | null | VIEW |



## TABLE - db_injection_backup 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_backup | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | data | db_injection_backup | null | NO | mediumtext | 16777215 | mediumtext |  | select,insert,update,references | 备份数据 |
| 3 | rollback_user | db_injection_backup |  | NO | varchar | 60 | varchar(60) |  | select,insert,update,references | 恢复执行人 |
| 4 | is_rolled_back | db_injection_backup | 0 | NO | tinyint | null | tinyint(4) |  | select,insert,update,references | 是否已恢复 |
| 5 | ct | db_injection_backup | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 创建时间 |
| 6 | rollback_time | db_injection_backup | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 回滚时间 |



## TABLE - db_injection_cluster 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_cluster | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | name | db_injection_cluster |  | NO | varchar | 160 | varchar(160) |  | select,insert,update,references | 名称 |
| 3 | description | db_injection_cluster |  | NO | varchar | 200 | varchar(200) |  | select,insert,update,references | 描述 |
| 4 | addr | db_injection_cluster |  | NO | varchar | 80 | varchar(80) |  | select,insert,update,references | 访问地址 |
| 5 | user | db_injection_cluster |  | NO | varchar | 120 | varchar(120) |  | select,insert,update,references | 用户 |
| 6 | pwd | db_injection_cluster |  | NO | varchar | 120 | varchar(120) |  | select,insert,update,references | 加密后秘钥 |
| 7 | ct | db_injection_cluster | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 创建时间 |
| 8 | ut | db_injection_cluster | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 更改时间 |
| 9 | operator | db_injection_cluster |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 操作人 |



## TABLE - db_injection_exec_item 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_exec_item | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | task_id | db_injection_exec_item | null | NO | bigint | null | bigint(20) unsigned |  | select,insert,update,references | task id |
| 3 | subtask_id | db_injection_exec_item | null | NO | bigint | null | bigint(20) unsigned |  | select,insert,update,references | subtask id |
| 4 | sql_content | db_injection_exec_item | null | YES | text | 65535 | text |  | select,insert,update,references | sql |
| 5 | remark | db_injection_exec_item |  | NO | varchar | 800 | varchar(800) |  | select,insert,update,references | 备注 |
| 6 | affect_rows | db_injection_exec_item | 0 | NO | int | null | int(11) |  | select,insert,update,references | 影响行数 |
| 7 | rule_comments | db_injection_exec_item |  | NO | varchar | 80 | varchar(80) |  | select,insert,update,references | 规则审核结果 |
| 8 | status | db_injection_exec_item |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 子项状态 |
| 9 | exec_info | db_injection_exec_item |  | NO | varchar | 200 | varchar(200) |  | select,insert,update,references | 执行信息 |
| 10 | backup_status | db_injection_exec_item |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 子项备份状态 |
| 11 | backup_info | db_injection_exec_item |  | NO | varchar | 200 | varchar(200) |  | select,insert,update,references | 备份异常信息 |
| 12 | backup_id | db_injection_exec_item | null | NO | bigint | null | bigint(20) unsigned |  | select,insert,update,references | backup id |
| 13 | ut | db_injection_exec_item | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 更改时间 |
| 14 | et | db_injection_exec_item | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 执行时间 |



## TABLE - db_injection_file 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_file | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | name | db_injection_file |  | NO | varchar | 80 | varchar(80) |  | select,insert,update,references | 文件名称 |
| 3 | location | db_injection_file |  | NO | varchar | 800 | varchar(800) |  | select,insert,update,references | 文件位置 |
| 4 | ct | db_injection_file | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 创建时间 |



## TABLE - db_injection_rule 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_rule | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | item | db_injection_rule |  | NO | varchar | 80 | varchar(80) |  | select,insert,update,references | 规则代号 |
| 3 | name | db_injection_rule | null | YES | varchar | 800 | varchar(800) |  | select,insert,update,references | 规则名称 |
| 4 | summary | db_injection_rule |  | NO | varchar | 800 | varchar(800) |  | select,insert,update,references | 规则摘要 |
| 5 | content | db_injection_rule | null | YES | varchar | 800 | varchar(800) |  | select,insert,update,references | 规则解释 |
| 6 | close | db_injection_rule | 0 | NO | int | null | int(11) |  | select,insert,update,references | 规则开关 |
| 7 | level | db_injection_rule | 0 | NO | int | null | int(11) |  | select,insert,update,references | 规则告警等级 |
| 8 | case | db_injection_rule |  | YES | varchar | 800 | varchar(800) |  | select,insert,update,references | 规则样例 |
| 9 | position | db_injection_rule | null | YES | int | null | int(11) |  | select,insert,update,references | 字符位置 |
| 10 | checkfunc | db_injection_rule |  | NO | varchar | 200 | varchar(200) |  | select,insert,update,references | 规则函数名 |
| 11 | threshold | db_injection_rule | null | YES | varchar | 200 | varchar(200) |  | select,insert,update,references | 阈值 |



## TABLE - db_injection_rule_status 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_rule_status | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | name | db_injection_rule_status |  | NO | varchar | 80 | varchar(80) |  | select,insert,update,references | 规则名 |
| 3 | close | db_injection_rule_status | 0 | NO | tinyint | null | tinyint(4) |  | select,insert,update,references | 开关 |



## TABLE - db_injection_subtask 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_subtask | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | task_id | db_injection_subtask | null | NO | bigint | null | bigint(20) unsigned |  | select,insert,update,references | task id |
| 3 | task_type | db_injection_subtask |  | NO | varchar | 10 | varchar(10) |  | select,insert,update,references | 任务类型 |
| 4 | db_name | db_injection_subtask |  | NO | varchar | 120 | varchar(120) |  | select,insert,update,references | 数据库名称 |
| 5 | cluster_name | db_injection_subtask |  | NO | varchar | 120 | varchar(120) |  | select,insert,update,references | 数据库集群 |



## TABLE - db_injection_task 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | db_injection_task | null | NO | bigint | null | bigint(20) unsigned | auto_increment | select,insert,update,references | 自增主键 |
| 2 | name | db_injection_task |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 名称 |
| 3 | status | db_injection_task |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 状态 |
| 4 | creator | db_injection_task |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 创建者 |
| 5 | reviewer | db_injection_task |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 审查者 |
| 6 | executor | db_injection_task |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 执行者 |
| 7 | exec_info | db_injection_task |  | NO | varchar | 40 | varchar(40) |  | select,insert,update,references | 执行信息 |
| 8 | reject_content | db_injection_task |  | NO | varchar | 500 | varchar(500) |  | select,insert,update,references | 驳回信息 |
| 9 | ct | db_injection_task | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 创建时间 |
| 10 | ut | db_injection_task | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 更改时间 |
| 11 | et | db_injection_task | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 执行时间 |
| 12 | ft | db_injection_task | 0 | NO | bigint | null | bigint(20) |  | select,insert,update,references | 执行结束时间 |



## TABLE - projectfile 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | projectfile | null | NO | int | null | int(11) | auto_increment | select,insert,update,references | 附件id |
| 2 | fileuploadercode | projectfile | null | YES | varchar | 128 | varchar(128) |  | select,insert,update,references | 附件上传者code |
| 3 | projectid | projectfile | null | YES | int | null | int(11) |  | select,insert,update,references | 项目id;此列受project表中的id列约束 |
| 4 | filename | projectfile | null | YES | varchar | 512 | varchar(512) |  | select,insert,update,references | 附件名 |
| 5 | fileurl | projectfile | null | YES | varchar | 512 | varchar(512) |  | select,insert,update,references | 附件下载地址 |
| 6 | filesize | projectfile | null | YES | bigint | null | bigint(20) |  | select,insert,update,references | 附件大小，单位Byte |



## TABLE - view2 
 表 dbinjection 中列信息
| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | view2 | 0 | NO | int | null | int(11) |  | select,insert,update,references | 附件id |
| 2 | fileuploadercode | view2 | null | YES | varchar | 128 | varchar(128) |  | select,insert,update,references | 附件上传者code |
| 3 | projectid | view2 | null | YES | int | null | int(11) |  | select,insert,update,references | 项目id;此列受project表中的id列约束 |
| 4 | filename | view2 | null | YES | varchar | 512 | varchar(512) |  | select,insert,update,references | 附件名 |
| 5 | fileurl | view2 | null | YES | varchar | 512 | varchar(512) |  | select,insert,update,references | 附件下载地址 |
| 6 | filesize | view2 | null | YES | bigint | null | bigint(20) |  | select,insert,update,references | 附件大小，单位Byte |


# DATABASE - test 
### 库 test 中表信息
| 序号 | 名称 | 来源 | 类型 | 存储引擎 | 创建时间 | 修改时间 | 备注 | 
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 0 |  |  |  |  |  |  |  |


