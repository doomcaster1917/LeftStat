CREATE TABLE IF NOT EXISTS view (
                                    id SERIAL PRIMARY KEY,
                                    name VARCHAR(250),
                                    title VARCHAR(250),
                                    img_addr VARCHAR(200),
                                    seo_description TEXT,
                                    seo_keywords TEXT,
                                    description TEXT
);
CREATE TABLE IF NOT EXISTS chart (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(250),
                                     title VARCHAR(250),
                                     description TEXT
);
CREATE TABLE IF NOT EXISTS dataset (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(250),
                                       data JSON, raw_data TEXT
);
CREATE TABLE IF NOT EXISTS dataset_chart (
                                             dataset_id INTEGER REFERENCES dataset (id)
                                                 ON UPDATE CASCADE ON DELETE CASCADE,
                                             chart_id INTEGER REFERENCES chart (id)
                                                 ON UPDATE CASCADE ON DELETE CASCADE,
                                             CONSTRAINT dataset_chart_pkey PRIMARY KEY (dataset_id, chart_id)
);
CREATE TABLE IF NOT EXISTS chart_view (
                                          chart_id INTEGER REFERENCES chart (id) ON UPDATE CASCADE,
                                          view_id INTEGER REFERENCES view (id) ON UPDATE CASCADE,
                                          CONSTRAINT chart_view_pkey PRIMARY KEY (chart_id, view_id)
);