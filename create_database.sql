create user pocockn;
alter user pocockn superuser;
drop database crypto_compare;
create database crypto_compare;
grant all privileges on database crypto_compare to pocockn;

drop database crypto_compare_test;
create database crypto_compare_test;
grant all privileges on database crypto_compare_test to pocockn;