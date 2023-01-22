--
-- PostgreSQL database dump
--

-- Dumped from database version 15.0
-- Dumped by pg_dump version 15.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: access_logs; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.access_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    personel_name text,
    personel_id_number text,
    lock_id bigint,
    key_id bigint,
    "timestamp" timestamp with time zone,
    location text
);


ALTER TABLE public.access_logs OWNER TO root;

--
-- Name: access_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.access_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.access_logs_id_seq OWNER TO root;

--
-- Name: access_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.access_logs_id_seq OWNED BY public.access_logs.id;


--
-- Name: access_rules; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.access_rules (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    personel_id bigint,
    lock_id bigint,
    key_id bigint,
    creator_id bigint,
    starts_at timestamp with time zone,
    ends_at timestamp with time zone
);


ALTER TABLE public.access_rules OWNER TO root;

--
-- Name: access_rules_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.access_rules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.access_rules_id_seq OWNER TO root;

--
-- Name: access_rules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.access_rules_id_seq OWNED BY public.access_rules.id;


--
-- Name: healthcheck_logs; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.healthcheck_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    lock_id bigint,
    "timestamp" timestamp with time zone,
    status boolean
);


ALTER TABLE public.healthcheck_logs OWNER TO root;

--
-- Name: healthcheck_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.healthcheck_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.healthcheck_logs_id_seq OWNER TO root;

--
-- Name: healthcheck_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.healthcheck_logs_id_seq OWNED BY public.healthcheck_logs.id;


--
-- Name: keys; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.keys (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    key_id text,
    label text,
    status boolean,
    description text
);


ALTER TABLE public.keys OWNER TO root;

--
-- Name: keys_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.keys_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.keys_id_seq OWNER TO root;

--
-- Name: keys_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.keys_id_seq OWNED BY public.keys.id;


--
-- Name: locks; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.locks (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    lock_id text,
    ip_address text,
    label text,
    description text,
    public_key text,
    location text
);


ALTER TABLE public.locks OWNER TO root;

--
-- Name: locks_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.locks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.locks_id_seq OWNER TO root;

--
-- Name: locks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.locks_id_seq OWNED BY public.locks.id;


--
-- Name: my_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.my_table (
    id bigint NOT NULL,
    name "char"[] NOT NULL,
    value "char"[] NOT NULL
);


ALTER TABLE public.my_table OWNER TO postgres;

--
-- Name: TABLE my_table; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.my_table IS 'buat nyobain bikin tabel aja';


--
-- Name: personels; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.personels (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    id_number text,
    name text,
    status boolean,
    role_id bigint,
    description text,
    key_id bigint
);


ALTER TABLE public.personels OWNER TO root;

--
-- Name: personels_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.personels_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.personels_id_seq OWNER TO root;

--
-- Name: personels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.personels_id_seq OWNED BY public.personels.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text
);


ALTER TABLE public.roles OWNER TO root;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO root;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: rssi_logs; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.rssi_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    rssi bigint,
    lock_id bigint,
    key_id bigint,
    "timestamp" timestamp with time zone
);


ALTER TABLE public.rssi_logs OWNER TO root;

--
-- Name: rssi_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.rssi_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rssi_logs_id_seq OWNER TO root;

--
-- Name: rssi_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.rssi_logs_id_seq OWNED BY public.rssi_logs.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text,
    username text,
    password text
);


ALTER TABLE public.users OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: access_logs id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.access_logs ALTER COLUMN id SET DEFAULT nextval('public.access_logs_id_seq'::regclass);


--
-- Name: access_rules id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.access_rules ALTER COLUMN id SET DEFAULT nextval('public.access_rules_id_seq'::regclass);


--
-- Name: healthcheck_logs id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.healthcheck_logs ALTER COLUMN id SET DEFAULT nextval('public.healthcheck_logs_id_seq'::regclass);


--
-- Name: keys id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.keys ALTER COLUMN id SET DEFAULT nextval('public.keys_id_seq'::regclass);


--
-- Name: locks id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.locks ALTER COLUMN id SET DEFAULT nextval('public.locks_id_seq'::regclass);


--
-- Name: personels id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.personels ALTER COLUMN id SET DEFAULT nextval('public.personels_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: rssi_logs id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.rssi_logs ALTER COLUMN id SET DEFAULT nextval('public.rssi_logs_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: access_logs; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.access_logs (id, created_at, updated_at, personel_name, personel_id_number, lock_id, key_id, "timestamp", location) FROM stdin;
3	2023-01-21 19:20:38.754634+07	2023-01-21 19:20:38.754634+07	ferb	1	1	1	2023-01-21 19:20:38.754634+07	\N
2	2023-01-21 19:20:38.754634+07	2023-01-21 19:20:38.754634+07	ferb	1	1	1	2023-01-21 19:20:38.754634+07	\N
4	2023-01-21 19:20:38.754634+07	2023-01-21 19:20:38.754634+07	ferb	1	1	1	2023-01-21 19:20:38.754634+07	\N
1	2023-01-21 19:20:38.754634+07	2023-01-21 19:20:38.754634+07	phineas	1	1	1	2023-01-21 19:20:38.754634+07	\N
\.


--
-- Data for Name: access_rules; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.access_rules (id, created_at, updated_at, personel_id, lock_id, key_id, creator_id, starts_at, ends_at) FROM stdin;
\.


--
-- Data for Name: healthcheck_logs; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.healthcheck_logs (id, created_at, updated_at, lock_id, "timestamp", status) FROM stdin;
1	2023-01-22 00:36:25.180508+07	2023-01-22 00:36:25.180508+07	1	2023-01-22 00:36:25.0663+07	f
2	2023-01-22 00:36:36.412577+07	2023-01-22 00:36:36.412577+07	1	2023-01-22 00:36:36.299969+07	f
3	2023-01-22 00:37:06.711988+07	2023-01-22 00:37:06.711988+07	1	2023-01-22 00:37:06.522852+07	f
4	2023-01-22 00:42:16.681306+07	2023-01-22 00:42:16.681306+07	1	2023-01-22 00:42:16.554606+07	f
5	2023-01-22 00:46:54.897141+07	2023-01-22 00:46:54.897141+07	5	2023-01-22 00:46:54.785082+07	f
6	2023-01-22 00:46:54.897141+07	2023-01-22 00:46:54.897141+07	1	2023-01-22 00:46:54.785082+07	f
7	2023-01-22 00:46:54.897141+07	2023-01-22 00:46:54.897141+07	3	2023-01-22 00:46:54.785082+07	f
8	2023-01-22 00:46:54.897141+07	2023-01-22 00:46:54.897141+07	2	2023-01-22 00:46:54.785082+07	f
9	2023-01-22 00:46:54.897141+07	2023-01-22 00:46:54.897141+07	4	2023-01-22 00:46:54.785082+07	f
10	2023-01-22 00:47:09.121627+07	2023-01-22 00:47:09.121627+07	5	2023-01-22 00:47:09.01054+07	f
11	2023-01-22 00:47:09.121627+07	2023-01-22 00:47:09.121627+07	2	2023-01-22 00:47:09.01054+07	f
12	2023-01-22 00:47:09.121627+07	2023-01-22 00:47:09.121627+07	1	2023-01-22 00:47:09.01054+07	f
13	2023-01-22 00:47:09.121627+07	2023-01-22 00:47:09.121627+07	3	2023-01-22 00:47:09.012052+07	f
14	2023-01-22 00:47:09.121627+07	2023-01-22 00:47:09.121627+07	4	2023-01-22 00:47:09.01739+07	f
\.


--
-- Data for Name: keys; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.keys (id, created_at, updated_at, key_id, label, status, description) FROM stdin;
1	2023-01-18 03:58:14.727942+07	2023-01-18 03:58:14.727942+07	DEADBEEFDEADBEEFDEADBEEFDEADBEEF	KEY-01	t	ini kunci KEY-01
2	2023-01-21 19:20:12.025039+07	2023-01-21 19:20:12.025039+07	DEADBEEFDEADBEEFDEADBEEFBEEFDEAD	KEY-02	t	ini kunci KEY-02
3	2023-01-21 19:20:26.157602+07	2023-01-21 19:20:26.157602+07	DEADBEEFDEADBEEFBEEFDEADDEADBEEF	KEY-03	t	ini kunci KEY-03
4	2023-01-21 19:20:38.754634+07	2023-01-21 19:20:38.754634+07	BEEFDEADBEEFDEADBEEFDEADDEADBEEF	KEY-04	t	ini kunci KEY-04
\.


--
-- Data for Name: locks; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.locks (id, created_at, updated_at, lock_id, ip_address, label, description, public_key, location) FROM stdin;
1	2023-01-18 03:54:20.780807+07	2023-01-18 03:54:20.780807+07	DEADBEEFDEADBEEFDEADBEEFDEADBEEF	127.0.0.1	Lock on 127.0.0.1		04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	
2	2023-01-18 03:54:20.780807+07	2023-01-18 03:54:20.780807+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE1	192.168.100.5	Lock on 192.168.100.5	\N	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	\N
3	2023-01-18 03:54:20.780807+07	2023-01-18 03:54:20.780807+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE2	192.168.100.6	Lock on 192.168.100.6	\N	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	\N
4	2023-01-18 03:54:20.780807+07	2023-01-18 03:54:20.780807+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE3	192.168.100.7	Lock on 192.168.100.7	\N	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	\N
5	2023-01-18 03:54:20.780807+07	2023-01-18 03:54:20.780807+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE4	192.168.100.8	Lock on 192.168.100.8	\N	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	\N
\.


--
-- Data for Name: my_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.my_table (id, name, value) FROM stdin;
\.


--
-- Data for Name: personels; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.personels (id, created_at, updated_at, id_number, name, status, role_id, description, key_id) FROM stdin;
1	2023-01-18 04:04:38.666893+07	2023-01-18 04:06:30.639425+07	22113344556677889900	phineas	f	1	adeknya phineas hehe	1
2	2023-01-18 04:05:13.981845+07	2023-01-18 04:05:13.981845+07	11223344556677889900	ferb	t	1		2
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.roles (id, created_at, updated_at, name) FROM stdin;
1	2023-01-18 04:04:06.466034+07	2023-01-18 04:04:06.466034+07	Guest
2	2023-01-18 04:04:06.474108+07	2023-01-18 04:04:06.474108+07	Doctor
3	2023-01-18 04:04:06.475792+07	2023-01-18 04:04:06.475792+07	Nurse
4	2023-01-18 04:04:06.477567+07	2023-01-18 04:04:06.477567+07	Staff
\.


--
-- Data for Name: rssi_logs; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.rssi_logs (id, created_at, updated_at, rssi, lock_id, key_id, "timestamp") FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.users (id, created_at, updated_at, name, username, password) FROM stdin;
1	2023-01-16 23:40:21.469672+07	2023-01-16 23:40:21.469672+07	admin baru	username	$2a$10$kWt1/3Avy5nQcmknQjgK3.5hkQWvI7eHGhd3lRd3BCvf4iVzbjDzq
2	2023-01-16 23:46:17.087108+07	2023-01-16 23:46:17.087108+07	admin baru lagi	username2	$2a$10$2ttRIFxC3Bec31w3jZYSzOsImClsdIqFVljkhInyXPRLUnIkp/IH6
\.


--
-- Name: access_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.access_logs_id_seq', 1, false);


--
-- Name: access_rules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.access_rules_id_seq', 1, false);


--
-- Name: healthcheck_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.healthcheck_logs_id_seq', 14, true);


--
-- Name: keys_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.keys_id_seq', 4, true);


--
-- Name: locks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.locks_id_seq', 2, true);


--
-- Name: personels_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.personels_id_seq', 2, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.roles_id_seq', 1, false);


--
-- Name: rssi_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.rssi_logs_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: access_logs access_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.access_logs
    ADD CONSTRAINT access_logs_pkey PRIMARY KEY (id);


--
-- Name: access_rules access_rules_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.access_rules
    ADD CONSTRAINT access_rules_pkey PRIMARY KEY (id);


--
-- Name: healthcheck_logs healthcheck_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.healthcheck_logs
    ADD CONSTRAINT healthcheck_logs_pkey PRIMARY KEY (id);


--
-- Name: personels idx_personels_id_number; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.personels
    ADD CONSTRAINT idx_personels_id_number UNIQUE (id_number);


--
-- Name: keys keys_key_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.keys
    ADD CONSTRAINT keys_key_id_key UNIQUE (key_id);


--
-- Name: keys keys_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.keys
    ADD CONSTRAINT keys_pkey PRIMARY KEY (id);


--
-- Name: locks locks_lock_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.locks
    ADD CONSTRAINT locks_lock_id_key UNIQUE (lock_id);


--
-- Name: locks locks_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.locks
    ADD CONSTRAINT locks_pkey PRIMARY KEY (id);


--
-- Name: my_table my_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.my_table
    ADD CONSTRAINT my_table_pkey PRIMARY KEY (id);


--
-- Name: personels personels_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.personels
    ADD CONSTRAINT personels_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: rssi_logs rssi_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.rssi_logs
    ADD CONSTRAINT rssi_logs_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- PostgreSQL database dump complete
--

