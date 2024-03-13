insert into usuarios (nome, nick, email, senha)
values
("Usuario 01", "Usuario 001", "Usuario001@gmail.com","$2a$10$RGAA3xhdlrxg4LEKEm7cmeNvGCd1dDi1FH.pTlRZBo0HY81./COka" ),
("Usuario 02", "Usuario 002", "Usuario002@gmail.com","$2a$10$RGAA3xhdlrxg4LEKEm7cmeNvGCd1dDi1FH.pTlRZBo0HY81./COka" ),
("Usuario 03", "Usuario 003", "Usuario003@gmail.com","$2a$10$RGAA3xhdlrxg4LEKEm7cmeNvGCd1dDi1FH.pTlRZBo0HY81./COka" );


insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("publicacoes do usuario 1", "essa é a publicaçao do usuario 1! ola ", 1),
("publicacoes do usuario 2", "essa é a publicaçao do usuario 2! ola ", 2),
("publicacoes do usuario 3", "essa é a publicaçao do usuario 3! ola ", 3);






