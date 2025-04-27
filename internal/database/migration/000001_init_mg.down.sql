-- Drop tables in reverse order of creation to handle foreign key dependencies
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS ticket_categories;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS users;