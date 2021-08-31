CREATE EXTENSION "uuid-ossp";

CREATE TYPE "exercise_category" AS ENUM (
  'HIIT',
  'Strength_Training',
  'Cardio',
  'Yoga',
  'Balance',
  'Flexibility',
  'Endurance',
  'Other',
  'Warm_Up',
  'Cool_Down'
);

CREATE TYPE "workout_split" AS ENUM (
  'Upper_Lower',
  'Full_Body',
  'Push_Pull_Leg',
  'Yoga',
  'Other'
);

CREATE TABLE "exercise" (
  "id" uuid DEFAULT uuid_generate_v1mc() PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "category" exercise_category,
  "name" text UNIQUE NOT NULL,
  "number_of_reps" int NOT NULL,
  "number_of_sets" int NOT NULL, 
  "repetition_unit" text,
  "photo_url" text,
  "video_url" text
);


CREATE TABLE "workout" (
  "id" uuid DEFAULT uuid_generate_v1mc() PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "name" text UNIQUE NOT NULL,
  "split" workout_split NOT NULL
);

CREATE TABLE "exercise_workout"(
  "id" uuid DEFAULT uuid_generate_v1mc() PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "exercise_id" UUID NOT NULL REFERENCES exercise(id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE CASCADE,
  "workout_id" UUID NOT NULL REFERENCES workout(id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE CASCADE
);

CREATE UNIQUE INDEX ON "exercise" ("id");

CREATE UNIQUE INDEX ON "workout" ("id");


