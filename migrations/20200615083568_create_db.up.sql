
CREATE TABLE IF NOT EXISTS "public"."matches" (
    "id" varchar NOT NULL,
    "matching_string" varchar NOT NULL,
    "name" varchar,
    "description" varchar,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "public"."logs" (
    "id" varchar NOT NULL,
    "log_line" varchar NOT NULL,
    "last_error" timestamp(0),
    "updated_at" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "error_count" int8,
    "time_start" int4,
    "time_end" int4,
    "matching_string" varchar,
    PRIMARY KEY ("id")
);
