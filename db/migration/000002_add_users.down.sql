ALTER table if exists "accounts" drop CONSTRAINT if exists "owner_currency_key";

ALTER table if exists "accounts" drop CONSTRAINT if exists "accounts_owner_fkey";

drop table if exists "users"