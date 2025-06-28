CREATE TABLE producto (
  descripcion varchar(100) DEFAULT NULL,
  precio decimal(15,2) NOT NULL,
  cantidad int(11) DEFAULT NULL,
  fecha date DEFAULT NULL,
  dolar decimal(6,2) NOT NULL,
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  bhabilitado tinyint(1) DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY id (id)
) 