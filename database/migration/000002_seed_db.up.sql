INSERT INTO "public"."exercise"("id", "category", "name", "number_of_reps", "number_of_sets", "repetition_unit", "photo_url", "video_url") 
	VALUES
	('fe232fb5-9325-439a-91bb-f73ee56374c9', 'Warm_Up', 'Walking knee hugs', 30, 1, 'seconds', NULL, NULL),
	('0ec2f4b5-8e86-49f6-9e07-3b8ebbd8209c','Warm_Up', 'Walking toe touches', 30, 1, 'seconds', NULL, NULL),
	('911fe8fe-23df-4bb7-b086-a2f24793b3b7','Strength_Training', 'Squat', 10, 3, 'count', NULL, NULL),
	('dabe3182-621b-4708-92ec-301d6a4c51ac','Strength_Training', 'Bridge', 10, 3, 'count', NULL, NULL),
	('fa4a3461-63c5-4bdc-bfc2-4dd5ac56315a','Strength_Training', 'Ball Exchange', 20, 3, 'count', NULL, NULL),
	('abc1e15c-1d08-451f-98d1-b889eb3f790c','Strength_Training', 'Plank to side plank', 10, 3, 'count', NULL, NULL),
	('7dff1226-ce40-492a-a752-02c1726982ba','Strength_Training', 'Kneeling chop', 12, 3, 'count', NULL, NULL),
	('72d051a7-419a-436f-b2b7-8782c3bb4393','Strength_Training', 'Press to Squat', 10, 3, 'count', NULL, NULL),
	('f4947cc8-44d9-4899-a894-65158e218792','Strength_Training', 'Chest Press', 10, 3, 'count', NULL, NULL),
	('2ebc7ec4-ca73-487c-b7ec-d7809b05d9ad','Strength_Training', 'Rows', 10, 3, 'count', NULL, NULL),
	('9a277974-cda9-4dd9-8d48-3d3dac2ef6a1','Strength_Training', 'Curls', 10, 3, 'count', NULL, NULL),
	('47659af0-b707-497d-9564-12522274a415','Strength_Training', 'Tricep pull downs', 10, 3, 'count', NULL, NULL),
	('92814ea2-b04a-4479-98dc-b0e81738a2f1','Cool_Down', 'Hip stretch', 30, 1, 'seconds', NULL, NULL),
	('fc0c3bb8-13d8-4a87-bc1a-4fa08dbb6cd8','Cool_Down', 'Child''s pose', 30, 1, 'seconds', NULL, NULL);
	
INSERT INTO "public"."workout"("id", "name", "split") 
	VALUES('cf89c692-050d-448c-ae57-4ba3bb4ee518', 'Workout 11/5', 'Full_Body');
	
INSERT INTO "public"."exercise_workout"("workout_id", "exercise_id")
	VALUES ('cf89c692-050d-448c-ae57-4ba3bb4ee518', 'fe232fb5-9325-439a-91bb-f73ee56374c9'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '0ec2f4b5-8e86-49f6-9e07-3b8ebbd8209c'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '911fe8fe-23df-4bb7-b086-a2f24793b3b7'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', 'dabe3182-621b-4708-92ec-301d6a4c51ac'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', 'fa4a3461-63c5-4bdc-bfc2-4dd5ac56315a'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', 'abc1e15c-1d08-451f-98d1-b889eb3f790c'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '7dff1226-ce40-492a-a752-02c1726982ba'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '72d051a7-419a-436f-b2b7-8782c3bb4393'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', 'f4947cc8-44d9-4899-a894-65158e218792'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '2ebc7ec4-ca73-487c-b7ec-d7809b05d9ad'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '9a277974-cda9-4dd9-8d48-3d3dac2ef6a1'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '47659af0-b707-497d-9564-12522274a415'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', '92814ea2-b04a-4479-98dc-b0e81738a2f1'),
	       ('cf89c692-050d-448c-ae57-4ba3bb4ee518', 'fc0c3bb8-13d8-4a87-bc1a-4fa08dbb6cd8');


INSERT INTO "public"."exercise"("id", "category", "name", "number_of_reps", "number_of_sets", "repetition_unit", "photo_url", "video_url") 
    VALUES
    ('d769de5e-91b8-449e-a89f-099557e00f09', 'Warm_Up', 'Cross over leg stretch', 30, 1, 'seconds', NULL, NULL),
    ('ac0969ba-feb6-451d-9004-b773dc5d3cd1','Warm_Up', 'High Knee March', 30, 1, 'seconds', NULL, NULL),
    ('d995ed8e-d514-49eb-b861-55d4a0974999','Strength_Training', 'Squat and curls', 10, 3, 'count', NULL, NULL),
    ('401bebb1-1b9b-4198-a15e-60d7341339f4','Strength_Training', 'Walking reverse lunges', 10, 3, 'count', NULL, NULL),
    ('ac3d196a-24cc-4b9f-b3ef-5b00e7405e5f','Strength_Training', 'Chest Flys', 10, 3, 'count', NULL, NULL),
    ('239006f8-6e7f-42ae-b79e-7d922fac19ee','Strength_Training', 'Deadlifts', 10, 3, 'count', NULL, NULL),
    ('0e965c2e-79cc-47ec-b675-a8ff8ec98fd7','Strength_Training', 'Romanian deadlift', 10, 3, 'count', NULL, NULL),
    ('78d0bfa8-85e0-42a6-8332-b0e4199dfc52','Strength_Training', 'Shoulder Press', 10, 3, 'count', NULL, NULL),
    ('98a3f55a-55ce-4dc4-ae9e-ed86f7358482','Strength_Training', 'Pull ups', 10, 3, 'count', NULL, NULL),
    ('c0893cdb-ddb2-49f3-b0e1-9b9fef09d592','Strength_Training', 'Reverse Curls', 10, 3, 'count', NULL, NULL),
    ('e926aaa0-7718-483f-acc9-0ae8910dd050','Strength_Training', 'Skull crushers', 10, 3, 'count', NULL, NULL),
    ('8eb5d302-4f57-4621-b39f-4bd5c6ee56bb','Cool_Down', 'Toe touches', 30, 1, 'seconds', NULL, NULL);
    
INSERT INTO "public"."workout"("id", "name", "split") 
    VALUES('d7911054-5940-4d8e-88ef-2c99d58a38af', 'Workout 11/14', 'Full_Body');
    
INSERT INTO "public"."exercise_workout"("workout_id", "exercise_id")
    VALUES ('d7911054-5940-4d8e-88ef-2c99d58a38af', 'd769de5e-91b8-449e-a89f-099557e00f09'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', 'ac0969ba-feb6-451d-9004-b773dc5d3cd1'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', 'd995ed8e-d514-49eb-b861-55d4a0974999'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', '401bebb1-1b9b-4198-a15e-60d7341339f4'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', 'ac3d196a-24cc-4b9f-b3ef-5b00e7405e5f'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', '239006f8-6e7f-42ae-b79e-7d922fac19ee'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', '0e965c2e-79cc-47ec-b675-a8ff8ec98fd7'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', '78d0bfa8-85e0-42a6-8332-b0e4199dfc52'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', '98a3f55a-55ce-4dc4-ae9e-ed86f7358482'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', 'c0893cdb-ddb2-49f3-b0e1-9b9fef09d592'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', 'e926aaa0-7718-483f-acc9-0ae8910dd050'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', 'fa4a3461-63c5-4bdc-bfc2-4dd5ac56315a'),
           ('d7911054-5940-4d8e-88ef-2c99d58a38af', '92814ea2-b04a-4479-98dc-b0e81738a2f1');


INSERT INTO "public"."exercise"("id", "category", "name", "number_of_reps", "number_of_sets", "repetition_unit", "photo_url", "video_url") 
    VALUES
    ('686cfa26-9f79-483a-8578-7f88403af9dc', 'Warm_Up', 'TRX chest stretch', 30, 1, 'seconds', NULL, NULL),
    ('2d1c6383-dd7f-4b05-b60d-ed33c1495534','Warm_Up', 'TRX lunge and twist', 30, 1, 'seconds', NULL, NULL),
    ('60654ddf-7c75-41f7-b82d-ec1069081aa3','Strength_Training', 'One legged Romanian deadlift', 10, 3, 'count', NULL, NULL),
    ('071fd3ab-0a1b-48ab-bf2b-abe5a74305b1','Strength_Training', 'Curl to press to extension', 10, 3, 'count', NULL, NULL),
    ('cc43eded-9654-4550-b302-ce2d64757460','Strength_Training', 'One arm row', 10, 3, 'count', NULL, NULL),
    ('ade5f3ec-2857-44cc-9d00-79621e1b8815','Strength_Training', 'Curl ladder', 10, 3, 'count', NULL, NULL),
    ('bc6caa11-c677-4ca4-8598-0db5674327ce','HIIT', 'Mountain climbers', 40, 3, 'seconds', NULL, NULL),
    ('94e1b073-2c8b-4c59-86c3-96adb2bbaab6','HIIT', 'Ice Skater Jumps', 60, 3, 'seconds', NULL, NULL),
    ('37d7b61b-b45a-42d9-bc28-4f9a690503cc','HIIT', 'Jumping Squats', 60, 3, 'seconds', NULL, NULL),
    ('0ca47502-1bff-4cc1-8599-8a9f994389db','HIIT', 'Russian Twist', 30, 3, 'count', NULL, NULL),
    ('4d7b2f74-0511-4175-a8fd-2a6c6a03a393','Strength_Training', 'Left to Right push ups', 10, 3, 'count', NULL, NULL),
    ('9d6ccc72-bc93-4b73-af29-c5557d8a0073','Cool_Down', 'Seated hip stretch', 30, 1, 'seconds', NULL, NULL);
    
INSERT INTO "public"."workout"("id", "name", "split") 
    VALUES('802dcb77-df35-47b5-b5df-029d95f5e664', 'Workout 12/17', 'Full_Body');
    
INSERT INTO "public"."exercise_workout"("workout_id", "exercise_id")
    VALUES ('802dcb77-df35-47b5-b5df-029d95f5e664', '686cfa26-9f79-483a-8578-7f88403af9dc'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '2d1c6383-dd7f-4b05-b60d-ed33c1495534'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '60654ddf-7c75-41f7-b82d-ec1069081aa3'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '071fd3ab-0a1b-48ab-bf2b-abe5a74305b1'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', 'cc43eded-9654-4550-b302-ce2d64757460'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', 'ade5f3ec-2857-44cc-9d00-79621e1b8815'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', 'bc6caa11-c677-4ca4-8598-0db5674327ce'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '94e1b073-2c8b-4c59-86c3-96adb2bbaab6'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '37d7b61b-b45a-42d9-bc28-4f9a690503cc'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '0ca47502-1bff-4cc1-8599-8a9f994389db'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '4d7b2f74-0511-4175-a8fd-2a6c6a03a393'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', 'fa4a3461-63c5-4bdc-bfc2-4dd5ac56315a'),
           ('802dcb77-df35-47b5-b5df-029d95f5e664', '9d6ccc72-bc93-4b73-af29-c5557d8a0073');

INSERT INTO "public"."exercise"("id", "category", "name", "number_of_reps", "number_of_sets", "repetition_unit", "photo_url", "video_url") 
    VALUES
    ('51d10705-6c93-4385-9b1a-9c707c12c5e0', 'Warm_Up', 'Calf wall stretch', 30, 1, 'seconds', NULL, NULL),
    ('dff785c3-dc4b-4cef-8979-066911aa55c0','Warm_Up', 'Boxer shuffle', 30, 1, 'seconds', NULL, NULL),
    ('789cf48c-4fb1-4b0b-9cd6-b29d80ed1b68','Strength_Training', 'Walk out shoulder touch push ups', 20, 3, 'count', NULL, NULL),
    ('e3a1bb2f-6d12-4765-a2fc-418686b5bfd2','Strength_Training', 'Low plank to high plank', 10, 3, 'count', NULL, NULL),
    ('debf763b-203c-4605-900c-db57becd1db0','Strength_Training', 'One legged lunge pulses', 20, 3, 'count', NULL, NULL),
    ('0be6f27b-ba31-4897-a925-7cba1329da4d','Strength_Training', 'Lateral Raise', 10, 3, 'count', NULL, NULL),
    ('c094915b-3acd-4dd8-9603-8db7ee6b86db','Strength_Training', 'Wall Sit', 40, 3, 'seconds', NULL, NULL),
    ('49c7ac39-e34b-4185-8086-f871bc17ac49','Strength_Training', 'Face Pulls', 15, 3, 'count', NULL, NULL),
    ('7c54339e-1f6e-4bf5-8bf5-61d7064d80c8','HIIT', 'Box Jump left to right', 45, 3, 'seconds', NULL, NULL),
    ('76ba57fc-9a4e-4136-aa0f-e9456415fcf7','Strength_Training', 'Exercise Band Shuffle', 30, 3, 'count', NULL, NULL),
    ('2a512684-0fb5-4525-8359-3823b010617e','Strength_Training', 'Hammer Curls', 10, 3, 'count', NULL, NULL);
    
INSERT INTO "public"."workout"("id", "name", "split") 
    VALUES('24acdc57-6a55-4d92-b070-d1f49c0666e5', 'Workout 12/21', 'Full_Body');
    
INSERT INTO "public"."exercise_workout"("workout_id", "exercise_id")
    VALUES ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '51d10705-6c93-4385-9b1a-9c707c12c5e0'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', 'dff785c3-dc4b-4cef-8979-066911aa55c0'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '789cf48c-4fb1-4b0b-9cd6-b29d80ed1b68'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', 'e3a1bb2f-6d12-4765-a2fc-418686b5bfd2'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', 'debf763b-203c-4605-900c-db57becd1db0'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '0be6f27b-ba31-4897-a925-7cba1329da4d'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', 'c094915b-3acd-4dd8-9603-8db7ee6b86db'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '49c7ac39-e34b-4185-8086-f871bc17ac49'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '7c54339e-1f6e-4bf5-8bf5-61d7064d80c8'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '76ba57fc-9a4e-4136-aa0f-e9456415fcf7'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '2a512684-0fb5-4525-8359-3823b010617e'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', 'fa4a3461-63c5-4bdc-bfc2-4dd5ac56315a'),
           ('24acdc57-6a55-4d92-b070-d1f49c0666e5', '9d6ccc72-bc93-4b73-af29-c5557d8a0073');


INSERT INTO "public"."exercise"("id", "category", "name", "number_of_reps", "number_of_sets", "repetition_unit", "photo_url", "video_url") 
    VALUES
    ('f2bf5855-d018-4eea-a4b8-e40339881579','Strength_Training', 'TRX Chest Press', 20, 3, 'count', NULL, NULL),
    ('33e204d8-e97c-4711-9c96-45bfe19f0805','Strength_Training', 'TRX Single Leg Squat', 15, 3, 'count', NULL, NULL),
    ('5c8a90d9-822d-4327-af2d-84ebdc17d1d1','HIIT', 'TRX Mountain Climbers', 60, 3, 'seconds', NULL, NULL),
    ('fe99e6e2-4505-458e-8e85-03ea5c2c743f','HIIT', 'Burpees', 15, 3, 'count', NULL, NULL),
    ('17312495-3038-4beb-94a6-68088578780a','Strength_Training', 'TRX Squat Jump', 20, 3, 'count', NULL, NULL),
    ('e6a13f22-538b-4837-b06c-8654763a704f','Strength_Training', 'TRX Atomic Push Ups', 15, 3, 'count', NULL, NULL),
    ('614223cc-66ba-4e71-b5ab-51193fbf4689','HIIT', 'Jump Rope', 90, 3, 'seconds', NULL, NULL),
    ('76ba57fc-9a4e-4136-aa0f-e9456415fcf7','Strength_Training', 'TRX Overhead Triceps Extension', 15, 3, 'count', NULL, NULL),
    ('7f39e836-5b80-45be-94a4-dbf40c78a60f','Strength_Training', 'TRX Single Leg Lunge', 15, 3, 'count', NULL, NULL),
    ('5afc2436-6590-4fa2-9b0f-2e4dd5f6d4b6','Strength_Training', 'TRX Abdominal Pike', 15, 3, 'count', NULL, NULL),
    ('24b13f3a-f14f-4d8b-b432-0cf2f6b93f02','HIIT', 'Ski Jumps', 90, 3, 'seconds', NULL, NULL),
    ('d45a825e-0dcf-4c1d-b00b-8c9afb45b74f','Strength_Training', 'TRX Bicep Curl', 15, 3, 'count', NULL, NULL),
    ('3a08bae4-d8c8-4e41-a5a8-0adb0585d8fa','Strength_Training', 'TRX Hamstring Curl', 15, 3, 'count', NULL, NULL),
    ('500d863d-17ec-4c3b-be77-c6d930f6472f','Strength_Training', 'TRX Side Plank', 45, 3, 'seconds', NULL, NULL),
    ('7982a4a5-dbba-466e-a8c6-e65066bb8798','Strength_Training', 'Jumping Lunges', 60, 3, 'seconds', NULL, NULL);

INSERT INTO "public"."workout"("id", "name", "split") 
    VALUES('b13faecb-1a79-40d1-873a-5bc3f0bf7265', 'Workout 03/10', 'Full_Body');
    
INSERT INTO "public"."exercise_workout"("workout_id", "exercise_id")
    VALUES ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '51d10705-6c93-4385-9b1a-9c707c12c5e0'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', 'dff785c3-dc4b-4cef-8979-066911aa55c0'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', 'f2bf5855-d018-4eea-a4b8-e40339881579'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '33e204d8-e97c-4711-9c96-45bfe19f0805'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '5c8a90d9-822d-4327-af2d-84ebdc17d1d1'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', 'fe99e6e2-4505-458e-8e85-03ea5c2c743f'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '17312495-3038-4beb-94a6-68088578780a'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', 'e6a13f22-538b-4837-b06c-8654763a704f'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '614223cc-66ba-4e71-b5ab-51193fbf4689'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '76ba57fc-9a4e-4136-aa0f-e9456415fcf7'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '7f39e836-5b80-45be-94a4-dbf40c78a60f'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '5afc2436-6590-4fa2-9b0f-2e4dd5f6d4b6'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '24b13f3a-f14f-4d8b-b432-0cf2f6b93f02'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', 'd45a825e-0dcf-4c1d-b00b-8c9afb45b74f'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '3a08bae4-d8c8-4e41-a5a8-0adb0585d8fa'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '500d863d-17ec-4c3b-be77-c6d930f6472f'),
           ('b13faecb-1a79-40d1-873a-5bc3f0bf7265', '7982a4a5-dbba-466e-a8c6-e65066bb8798');
