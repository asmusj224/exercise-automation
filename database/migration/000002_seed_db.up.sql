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
           