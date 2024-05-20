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

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);