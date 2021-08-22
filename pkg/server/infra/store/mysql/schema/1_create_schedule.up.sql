# 予約枠管理テーブル(管理者の予約枠定義と、予約在庫の管理に利用)
CREATE TABLE IF NOT EXISTS schedule
(
    id              VARCHAR(255) CHARACTER SET ascii NOT NULL, # uid
    date            DATE                             NOT NULL,
    hour            TINYINT      UNSIGNED            NOT NULL,
    min             TINYINT      UNSIGNED            NOT NULL,
    max_available   INTEGER      UNSIGNED            NOT NULL,
    stock           INTEGER      UNSIGNED            NOT NULL,
    created_at      DATETIME                         NOT NULL,
    updated_at      DATETIME                         NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (date, hour, min)
);
