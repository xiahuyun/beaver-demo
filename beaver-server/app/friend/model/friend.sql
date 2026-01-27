CREATE TABLE friend (
    id bigint NOT NULL AUTO_INCREMENT,
    friendId varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The friend id',
    friendshipId varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The friendship id',
    sendUserId varchar(255) NOT NULL DEFAULT '' COMMENT 'The send user id',
    revUserId varchar(255) NOT NULL DEFAULT '' COMMENT 'The receive user id',
    sendUserNotice varchar(255) NOT NULL DEFAULT '' COMMENT 'The send user notice',
    revUserNotice varchar(255) NOT NULL DEFAULT '' COMMENT 'The receive user notice',
    create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    source varchar(255) NOT NULL DEFAULT '' COMMENT 'The friend source',
    isDeleted varchar(255) NOT NULL DEFAULT '' COMMENT 'The friend is deleted',
    UNIQUE friend_id_index (friendId),
    UNIQUE friendship_id_index (friendshipId),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'friend table';