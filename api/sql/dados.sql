insert into usuarios (nome, nick, email, senha)
values
("Usuáro 1", "usuario_1", "usuario1@gmail.com", "$2a$10$mysMxArZJBPUB4FP8ffB9urGe.MmIHe5UNhNsVW93X1mRugZ.HvY6"),
("Usuáro 2", "usuario_2", "usuario2@gmail.com", "$2a$10$mysMxArZJBPUB4FP8ffB9urGe.MmIHe5UNhNsVW93X1mRugZ.HvY6"),
("Usuáro 3", "usuario_3", "usuario3@gmail.com", "$2a$10$mysMxArZJBPUB4FP8ffB9urGe.MmIHe5UNhNsVW93X1mRugZ.HvY6");


insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);