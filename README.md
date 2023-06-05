# PULIDs (Prefixed ULIDs)

PULIDs are an identifier encoding scheme that builds upon the excellent [Universally Unique Lexicographically Sortable Identifier](https://github.com/oklog/ulid) scheme. Prefixes should maintain compatibility with base32. The library uses the following alphabet: `0123456789ABCDEFGHJKMNPQRSTVWXYZ`. If constrained to 2 characters for encoding the entity type, the number of types of entities that can be referenced with this scheme is 2^32 (1024). Though keep in mind that obvious encodings of table names will start to overlap far sooner.
