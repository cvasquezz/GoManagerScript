create database postgres
;

comment on database postgres is 'default administrative connection database'
;

create table if not exists job_users
(
		usr_id serial not null
				constraint job_users_pkey
							primary key,
								usr_nombre varchar(50),
									usr_ape varchar(70),
										usr_user varchar(30) not null,
											usr_fono varchar(20),
												usr_activo integer not null,
													usr_pass varchar(150) not null
												)
												;

												comment on table job_users is 'Tabla de usuarios para autenticacion al sistema'
												;

												create table if not exists job
												(
														job_cod varchar(100),
															job_mallid varchar(100),
																job_status varchar(100),
																	job_fec_ini timestamp,
																		job_fec_fin timestamp,
																			job_id serial not null
																					constraint job_job_id_pk
																								primary key
																							)
																							;

																							comment on table job is 'Tabla de semanforo, para vidar los estados de los procesos'
																							;

																							create unique index if not exists job_job_id_uindex
																								on job (job_id)
																								;

																								create table if not exists job_contacts
																								(
																										con_id serial not null
																												constraint job_contacts_pkey
																															primary key,
																																con_kind varchar(100),
																																	con_email varchar(100),
																																		con_status integer
																																	)
																																	;

																																	comment on table job_contacts is 'Tabla para almacenar los correos de los destinatarios por proceso'
																																	;

																																	create table if not exists process_script
																																	(
																																			pro_id serial not null
																																					constraint process_script_pkey
																																								primary key,
																																									pro_kind varchar(100),
																																										pro_path varchar(200),
																																											pro_script varchar(100),
																																												pro_lang varchar(50),
																																													pro_function varchar(100)
																																												)
																																												;

																																												comment on table process_script is 'tabla con los scripts para ejecutar controlados desde app'
																																												;

																																												create table if not exists log_process
																																												(
																																														log_id serial not null
																																																constraint log_process_pkey
																																																			primary key,
																																																				job_cod varchar(100),
																																																					job_mallid varchar(100),
																																																						log_pid varchar(20),
																																																							log_date timestamp,
																																																								log_texto varchar(250)
																																																							)
																																																							;

																																																							comment on table log_process is 'Tabla para capturar la salida estandar de los scripts'
																																																							;

																																																							create table if not exists log_sender
																																																							(
																																																									id serial not null
																																																											constraint log_sender_pkey
																																																														primary key,
																																																															pid varchar(20)
																																																														)
																																																														;

																																																														comment on table log_sender is 'Tabla pra almacenar los pid de los logs enviados a los destinatarios'
																																																														;


