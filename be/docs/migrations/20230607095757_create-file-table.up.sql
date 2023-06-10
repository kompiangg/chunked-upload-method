create table if not exists file (
  id int primary key auto_increment,
  name text,
  url text,
  created_at datetime default current_timestamp()
);