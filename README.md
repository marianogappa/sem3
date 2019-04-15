# sem3
Parses the Semantics3 category tree to TSV format (CSV using tab as separator). TSV allows you to copy-paste to Excel/Google Sheets, or upload to a DB.

## What does it do?

Only this:

```
$ sem3 | head
13157	Beauty
12122	Beauty	Tools & Accessories
20354	Beauty	Tools & Accessories	Hair Coloring Tools
7504	Beauty	Tools & Accessories	Hair Coloring Tools	Hair Color Removers
19186	Beauty	Tools & Accessories	Hair Coloring Tools	Brushes, Combs & Needles
7131	Beauty	Tools & Accessories	Hair Coloring Tools	Hair Color Mixing Bowls
14511	Beauty	Tools & Accessories	Hair Coloring Tools	Applicator Bottles
12483	Beauty	Tools & Accessories	Hair Coloring Tools	Caps, Foils & Wraps
21323	Beauty	Tools & Accessories	Nail Tools
10083	Beauty	Tools & Accessories	Nail Tools	Cuticle Pushers
```

## How do I use it?

```
$ sem3 > categories.csv
```

## What are the fields?

1. Category Id
2. Main category
3. Sub category
4. Sub-sub category
5. Sub-sub-sub category

3, 4 and 5 can be empty.

```
CREATE TABLE categories (
    id integer PRIMARY KEY,
    cat1 character varying(255),
    cat2 character varying(255),
    cat3 character varying(255),
    cat4 character varying(255)
);
```

## How do I upload to a MySQL table?

```
LOAD DATA INFILE "/path/to/categories.csv"
INTO TABLE categories
COLUMNS TERMINATED BY '\t'
LINES TERMINATED BY '\n'
```

## How do I upload to a Postgres table?

```
COPY categories FROM '/path/to/categories.csv' DELIMITERS E'\t' CSV;
```
