INSERT INTO public.job (job_cod, job_mallid, job_status, job_fec_ini, job_fec_fin, job_id) VALUES ('ejemplo', 'sh', 'IDLE', '2019-01-30 23:59:44.278269', '2019-01-31 00:00:09.358840', 1);
INSERT INTO public.job_contacts (con_id, con_kind, con_email, con_status) VALUES (1, 'ejemplo', 'christian.vasquez@karibu.cl', 1);
INSERT INTO public.job_users (usr_id, usr_nombre, usr_ape, usr_user, usr_fono, usr_activo, usr_pass) VALUES (1, 'Christian', 'Vásquez Zavala', 'cvasquez', '+56940103762', 1, 'e10adc3949ba59abbe56e057f20f883e');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (1, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.031972', 'mi pid es 7698');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (2, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.037856', 'incio de ejecucion  2019-01-30 23:58:30');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (3, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.042292', 'hola');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (4, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.047031', 'este es un ejemplo parametro uno = ejemplo parametro dos sh');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (5, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.052683', 'se agregan nuevas lineas de log para pruebas');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (6, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.056907', 'no se me ocurre que mas colocar XD');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (7, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.061287', 'chao');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (8, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.070417', 'Lorea don pato, esta es la salida estandar de bash capturada con golang y enviada por email con go');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (9, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.082812', '8) soy el puto amo xD');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (10, 'ejemplo', 'sh', '7698', '2019-01-30 23:58:30.087099', 'fin de ejecucion  2019-01-30 23:58:30');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (11, 'ejemplo', 'sh', '7846', '2019-01-30 23:59:44.287561', 'mi pid es 7846');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (12, 'ejemplo', 'sh', '7846', '2019-01-30 23:59:44.291956', 'incio de ejecucion  2019-01-30 23:59:44');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (13, 'ejemplo', 'sh', '7846', '2019-01-30 23:59:44.297249', 'hola');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (14, 'ejemplo', 'sh', '7846', '2019-01-30 23:59:49.304457', 'este es un ejemplo parametro uno = ejemplo parametro dos sh');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (15, 'ejemplo', 'sh', '7846', '2019-01-30 23:59:59.310688', 'se agregan nuevas lineas de log para pruebas');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (16, 'ejemplo', 'sh', '7846', '2019-01-30 23:59:59.323529', 'no se me ocurre que mas colocar XD');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (17, 'ejemplo', 'sh', '7846', '2019-01-30 23:59:59.335957', 'chao');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (18, 'ejemplo', 'sh', '7846', '2019-01-31 00:00:09.318005', 'Lorea don pato, esta es la salida estandar de bash capturada con golang y enviada por email con go');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (19, 'ejemplo', 'sh', '7846', '2019-01-31 00:00:09.331570', '8) soy el puto amo xD');
INSERT INTO public.log_process (log_id, job_cod, job_mallid, log_pid, log_date, log_texto) VALUES (20, 'ejemplo', 'sh', '7846', '2019-01-31 00:00:09.343740', 'fin de ejecucion  2019-01-31 00:00:09');
INSERT INTO public.log_sender (id, pid) VALUES (1, '');
INSERT INTO public.log_sender (id, pid) VALUES (2, '');
INSERT INTO public.log_sender (id, pid) VALUES (3, '');
INSERT INTO public.log_sender (id, pid) VALUES (4, '');
INSERT INTO public.log_sender (id, pid) VALUES (5, '');
INSERT INTO public.log_sender (id, pid) VALUES (6, '');
INSERT INTO public.log_sender (id, pid) VALUES (7, '');
INSERT INTO public.log_sender (id, pid) VALUES (8, '');
INSERT INTO public.log_sender (id, pid) VALUES (9, '');
INSERT INTO public.log_sender (id, pid) VALUES (10, '');
INSERT INTO public.log_sender (id, pid) VALUES (11, '');
INSERT INTO public.log_sender (id, pid) VALUES (12, '');
INSERT INTO public.log_sender (id, pid) VALUES (13, '');
INSERT INTO public.log_sender (id, pid) VALUES (14, '');
INSERT INTO public.log_sender (id, pid) VALUES (15, '');
INSERT INTO public.log_sender (id, pid) VALUES (16, '');
INSERT INTO public.log_sender (id, pid) VALUES (17, '');
INSERT INTO public.log_sender (id, pid) VALUES (18, '');
INSERT INTO public.log_sender (id, pid) VALUES (19, '');
INSERT INTO public.log_sender (id, pid) VALUES (20, '');
INSERT INTO public.log_sender (id, pid) VALUES (21, '');
INSERT INTO public.log_sender (id, pid) VALUES (22, '');
INSERT INTO public.log_sender (id, pid) VALUES (23, '');
INSERT INTO public.log_sender (id, pid) VALUES (24, '7698');
INSERT INTO public.log_sender (id, pid) VALUES (25, '7846');
INSERT INTO public.process_script (pro_id, pro_kind, pro_path, pro_script, pro_lang, pro_function) VALUES (1, 'ejemplo', '/home/christian/go/src/github.com/GoManagerScript', 'eje.sh', 'sh', null);