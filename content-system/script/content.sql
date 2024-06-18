CREATE DATABASE `cms_content` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE `t_content_details` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '内容ID',
    `title` varchar(255) DEFAULT '' COMMENT '内容标题',
    `description` varchar(500) DEFAULT '' COMMENT '内容描述',
    `author` varchar(64) DEFAULT '' COMMENT '作者',
    `video_url` varchar(255) DEFAULT '' COMMENT '视频链接',
    `thumbnail` varchar(255) DEFAULT '' COMMENT '封面图',
    `category` varchar(64) DEFAULT '' COMMENT '分类',
    `duration`  int DEFAULT 0 COMMENT '时长',
    `resolution` varchar(64) DEFAULT '' COMMENT '分辨率',
    `file_size` int DEFAULT 0 COMMENT '文件大小',
    `format` varchar(64) DEFAULT '' COMMENT '格式',
    `quality` int DEFAULT 0 COMMENT '视频质量 1-高清 2-标清 3-流畅',
    `approval_status` int DEFAULT 0 COMMENT '审核状态 0-未审核 1-审核中 2-审核通过 3-审核不通过',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='内容表';