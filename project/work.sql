CREATE TABLE `product`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `name`        varchar(64)         not null default '' comment '产品名',
    `product_key` bigint(20) unsigned not null default 0 comment '产品ID',
    `version_id`  bigint(20) unsigned not null default 0 comment '版本ID',
    `processes`   text comment '所需工序',
    `note`        text comment '备注',
    `extra`       text comment '其他信息',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`),
    unique product_version (`product_key`, `version_id`)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='产品表';

CREATE TABLE `ticket`
(
    `id`              bigint(20) unsigned not null auto_increment comment '自增ID',
    `name`            varchar(64)         not null default '' comment '工单名',
    `product_id`      bigint(20) unsigned not null default 0 comment '对应产品ID',
    `product_key`     bigint(20) unsigned not null default 0 comment '对应产品key',
    `product_version` bigint(20) unsigned not null default 0 comment '对应产品版本',
    `ticket_key`      bigint(20) unsigned not null default 0 comment '工单key',
    `product_count`   int                 not null default 0 comment '产品数量',
    `note`            text comment '备注',
    `extra`           text comment '其他信息',
    `create_time`     timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`     timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='工单表';

CREATE TABLE `test_result`
(
    `id`                     bigint(20) unsigned not null auto_increment comment '自增ID',
    `ticket_id`              bigint(20) unsigned not null default 0 comment '工单号',
    `node_key`               bigint(20) unsigned not null default 0 comment '节点key',
    `node_name`              varchar(64)         not null default '' comment '节点名',
    `node_num`               bigint(20) unsigned not null default 0 comment '节点序号',
    `node_requirement_upper` decimal(10, 5)      not null default 0 comment '要求上界',
    `node_requirement_lower` decimal(10, 5)      not null default 0 comment '要求下界',
    `node_requirement_unit`  varchar(32)         not null default '' comment '要求单位',
    `node_test_value`        decimal(10, 5)      not null default 0 comment '实际测量值',
    `tester`                 bigint(20)          not null default 0 comment '检测者',
    `note`                   text comment '备注',
    `extra`                  text comment '其他信息',
    `create_time`            timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`            timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='检测数据表';

CREATE TABLE `tester`
(
    `id`          bigint(20) unsigned not null auto_increment comment '自增ID',
    `name`        varchar(64)         not null default '' comment '名字',
    `note`        text comment '备注',
    `extra`       text comment '其他信息',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='人员表';