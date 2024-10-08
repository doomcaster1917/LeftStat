create type chart_type as enum ('line', 'bar', 'pie');
ALTER TABLE chart ADD chart_type chart_type DEFAULT('line');