create user crypto_user with password 'password';
alter user crypto_user superuser;
drop database crypto_compare;
create database crypto_compare;
grant all privileges on database crypto_compare to crypto_user;

drop database crypto_compare_test;
create database crypto_compare_test;
grant all privileges on database crypto_compare_test to crypto_user;