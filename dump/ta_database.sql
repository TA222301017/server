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
    location text,
    status boolean
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
    "timestamp" timestamp with time zone,
    personel_id bigint
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
5	2023-02-04 20:59:42.423114+07	2023-02-04 20:59:42.423114+07	Ipin Saipudin	22113344556677889900	26	1	2023-02-04 20:59:42.422565+07	
6	2023-02-04 20:59:50.888218+07	2023-02-04 20:59:50.888218+07	Ipin Saipudin	22113344556677889900	26	1	2023-02-04 20:59:50.888218+07	
7	2023-02-04 20:59:55.264279+07	2023-02-04 20:59:55.264279+07	Ipin Saipudin	22113344556677889900	26	1	2023-02-04 20:59:55.263892+07	
8	2023-02-04 21:13:06.071496+07	2023-02-04 21:13:06.071496+07	Muhamad Taruna	13219030	26	3	2023-02-04 21:13:06.071235+07	
\.


--
-- Data for Name: access_rules; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.access_rules (id, created_at, updated_at, personel_id, lock_id, key_id, creator_id, starts_at, ends_at) FROM stdin;
42	2023-01-24 01:02:11.291843+07	2023-01-24 01:02:18.911936+07	1	2	0	1	2023-01-24 00:00:00+07	2023-01-24 23:59:00+07
45	2023-02-04 22:11:25.67281+07	2023-02-04 22:11:25.67281+07	2	1	0	1	2023-02-04 00:00:00+07	2023-02-04 23:59:00+07
46	2023-02-04 22:11:31.576385+07	2023-02-04 22:11:31.576385+07	2	22	0	1	2023-02-04 00:00:00+07	2023-02-04 23:59:00+07
54	2023-02-05 14:50:20.471445+07	2023-02-05 16:50:32.224416+07	4	1	3	1	2023-02-01 03:20:00+07	2023-02-05 00:00:00+07
36	2023-01-23 21:08:51.255956+07	2023-01-23 21:08:51.255956+07	5	1	0	1	2023-01-23 00:00:00+07	2023-01-23 23:59:00+07
37	2023-01-23 21:08:58.043611+07	2023-01-24 00:38:52.370789+07	5	1	0	1	2023-01-24 00:00:00+07	2023-01-24 23:59:00+07
38	2023-01-24 00:08:42.546957+07	2023-01-24 00:42:34.046454+07	1	5	0	1	2023-01-24 00:00:00+07	2023-01-24 23:59:00+07
39	2023-01-24 00:43:22.891002+07	2023-01-24 00:43:22.891002+07	2	4	0	1	2023-01-24 00:00:00+07	2023-01-24 23:59:00+07
41	2023-01-24 01:00:07.731833+07	2023-01-24 01:00:07.731833+07	4	5	0	1	2023-01-24 00:00:00+07	2023-01-24 23:59:00+07
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
15	2023-02-03 21:08:08.613486+07	2023-02-03 21:08:08.613486+07	4	2023-02-03 21:08:08.51174+07	f
16	2023-02-03 21:08:21.584328+07	2023-02-03 21:08:21.584328+07	2	2023-02-03 21:08:21.474061+07	f
17	2023-02-03 21:11:52.863685+07	2023-02-03 21:11:52.863685+07	5	2023-02-03 21:11:52.753061+07	f
18	2023-02-03 21:18:24.134959+07	2023-02-03 21:18:24.134959+07	5	2023-02-03 21:18:24.013407+07	f
19	2023-02-05 17:08:22.643951+07	2023-02-05 17:08:22.643951+07	19	2023-02-05 17:08:22.519855+07	f
20	2023-02-05 17:08:31.133327+07	2023-02-05 17:08:31.133327+07	21	2023-02-05 17:08:31.017862+07	f
21	2023-02-05 17:08:33.908869+07	2023-02-05 17:08:33.908869+07	22	2023-02-05 17:08:33.791397+07	f
22	2023-02-05 17:09:34.114924+07	2023-02-05 17:09:34.114924+07	22	2023-02-05 17:09:34.108234+07	f
23	2023-02-05 17:11:32.050994+07	2023-02-05 17:11:32.050994+07	26	2023-02-05 17:11:31.937356+07	f
24	2023-02-05 17:15:38.255236+07	2023-02-05 17:15:38.255236+07	26	2023-02-05 17:15:38.243564+07	f
25	2023-02-05 17:16:03.971338+07	2023-02-05 17:16:03.971338+07	26	2023-02-05 17:16:03.95836+07	f
26	2023-02-05 17:16:13.76017+07	2023-02-05 17:16:13.76017+07	26	2023-02-05 17:16:13.754985+07	f
27	2023-02-05 17:22:41.723438+07	2023-02-05 17:22:41.723438+07	26	2023-02-05 17:22:41.71704+07	t
28	2023-02-05 17:22:48.84673+07	2023-02-05 17:22:48.84673+07	26	2023-02-05 17:22:48.842425+07	t
\.


--
-- Data for Name: keys; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.keys (id, created_at, updated_at, key_id, label, status, description) FROM stdin;
2	2023-01-21 19:20:12.025039+07	2023-01-21 19:20:12.025039+07	DEADBEEFDEADBEEFDEADBEEFBEEFDEAD	KEY-02	t	ini kunci KEY-02
3	2023-01-21 19:20:26.157602+07	2023-01-21 19:20:26.157602+07	DEADBEEFDEADBEEFBEEFDEADDEADBEEF	KEY-03	t	ini kunci KEY-03
5	2023-01-22 16:37:23.618966+07	2023-01-22 16:37:23.618966+07	DEADBEEFDEADBEEFDEADBEEFBEEFDDDD	KUNCI-MATI	t	kunci ini mati
6	2023-01-22 21:18:19.007361+07	2023-01-22 21:18:19.007361+07	ASDFASDFASDFASDFASDFASDFASDFASDF	asdf	t	asdf
8	2023-01-22 21:22:31.102778+07	2023-01-22 21:22:31.102778+07	FDSAFDSAFDSAFDSAFDSAFDSAFDSAFDSA	FDSA	t	FDSA
13	2023-01-22 21:23:18.480572+07	2023-01-22 21:23:18.480572+07	FDSAFDSAFDSAFDSAFDSAFDSAFDSAFDSB	FDSA	t	FDSA
1	2023-01-18 03:58:14.727942+07	2023-01-22 21:48:05.890721+07	DEADBEEFDEADBEEFDEADBEEFDEADBEEF	KEY-01	t	Ini kunci mantap lah pokoknya
14	2023-01-23 00:02:15.989935+07	2023-01-23 00:02:15.989935+07	FDSAFDSAFDSAFDSAFDSAFDSAFDSAFDSC	FDSC	t	FDSC
4	2023-01-21 19:20:38.754634+07	2023-02-02 23:04:19.416009+07	BEEFDEADBEEFDEADBEEFDEADDEADBEEF	KEY-04	t	ini kunci KEY-04
\.


--
-- Data for Name: locks; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.locks (id, created_at, updated_at, lock_id, ip_address, label, description, public_key, location, status) FROM stdin;
20	2023-02-04 19:55:50.431737+07	2023-02-04 19:55:50.431737+07	BBD8AC2EC653ED8A23038AF3EFA88F3E	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
23	2023-02-04 20:01:57.704759+07	2023-02-04 20:01:57.704759+07	0C5346870A2DD5E57FD13F0F9F8F63D3	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
24	2023-02-04 20:01:57.721005+07	2023-02-04 20:01:57.721005+07	612A719B74F5646B9384B68D00DE20C4	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
1	2023-01-18 03:54:20.780807+07	2023-01-22 21:33:22.291357+07	DEADBEEFDEADBEEFDEADBEEFDEADBEEF	127.0.0.1	Lock on 127.0.0.1	Lorem ipsum dolor sit amet	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	Sayap 1, Lantai 3, Gedung Asep	f
2	2023-01-18 03:54:20.780807+07	2023-01-24 01:13:45.392661+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE1	192.168.100.5	Lock on 192.168.100.5	Lock di Gudang A	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	Gudang A	f
3	2023-01-18 03:54:20.780807+07	2023-01-18 03:54:20.780807+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE2	192.168.100.6	Lock on 192.168.100.6	\N	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	\N	t
4	2023-01-18 03:54:20.780807+07	2023-01-18 03:54:20.780807+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE3	192.168.100.7	Lock on 192.168.100.7	\N	04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d	\N	t
25	2023-02-04 20:46:06.819411+07	2023-02-04 20:46:06.819411+07	E0ACF5507BB6C14CBB23220CE9D3530E	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
5	2023-01-18 03:54:20.780807+07	2023-02-03 21:18:24.124755+07	DEADBEEFDEADBEEFDEADBEEFDEADBEE4	192.168.100.8	Lock on 192.168.100.8		04348766ba5b36436fd3e2528cf0a979bdc891f3bde663438377e99e70d6428d3ed26d387e1e35c5395a622402ed63f2ca2306f4eeae8a9c8d		f
6	2023-02-04 19:04:30.485389+07	2023-02-04 19:04:30.485389+07	04EEC58A8EA4E6A6A536054ACC4DCC90	127.0.0.1	Lock on 127.0.0.1		5f88bc3a1af219da38594ddc446ae4bf859e097139d3eb1164411e3811c70080ec561f442d5b4e2f33		f
11	2023-02-04 19:46:36.917402+07	2023-02-04 19:46:36.917402+07	21F393BA3811C1CA62B68E125F843CBB	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
12	2023-02-04 19:46:37.140692+07	2023-02-04 19:46:37.140692+07	C83AEFEDBB61B355E9A84D5221793645	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
13	2023-02-04 19:46:54.847262+07	2023-02-04 19:46:54.847262+07	632808B871BAE76AA36A96B4D96A730E	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
14	2023-02-04 19:46:54.907654+07	2023-02-04 19:46:54.907654+07	FB777214F43CA730B587A35B1B4E6203	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
15	2023-02-04 19:48:30.371648+07	2023-02-04 19:48:30.371648+07	EEB7D4B34333E4599366C6A20178EF1A	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
16	2023-02-04 19:48:30.409735+07	2023-02-04 19:48:30.409735+07	BBA3A01A052B20D0718B1908755D810A	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
17	2023-02-04 19:52:55.350484+07	2023-02-04 19:52:55.350484+07	4E6737A7CFD2090EFDB8263441CECC4E	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
18	2023-02-04 19:52:55.395866+07	2023-02-04 19:52:55.395866+07	AC4FC5951141B9650CDA9DA48586E3B7	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
19	2023-02-04 19:55:50.378567+07	2023-02-05 17:08:22.630172+07	2E6ADE0E600002B5DB6BD105B9150E8D	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
21	2023-02-04 19:59:34.849524+07	2023-02-05 17:08:31.127996+07	CE5752782DB482B1E9FFC54ABDDF83C2	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
22	2023-02-04 19:59:34.863433+07	2023-02-05 17:09:34.108234+07	092C140948845F2771C16F7E4BCA7F29	127.0.0.1	Lock on 127.0.0.1		04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		f
26	2023-02-04 20:47:23.793271+07	2023-02-05 17:24:04.426017+07	A006BA706CFBCC535E5C67A23DDB45B8	127.0.0.1	Mock Lock	Mock dari subsistem lock yang digunakan untuk keperluan coba2	04326fdbac250aa84617a6b55a0ee659f32959aca41e9ea9e815fc07c640e7e727eff3b072ff0d19fdf8ee13285f4998d49db11f6fa3b27f11		t
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
2	2023-01-18 04:05:13.981845+07	2023-01-23 04:24:27.326519+07	11223344556677889900	Upin Odinson	t	3	Anaknya pak Odin	4
3	2023-01-23 04:31:41.402778+07	2023-01-23 04:31:41.402778+07	13219059	Elkhan Julian Brillianshah	t	4	Anggota TA222301017	8
4	2023-01-23 04:32:33.253064+07	2023-01-23 04:32:33.253064+07	13219030	Muhamad Taruna	t	4	Anggota TA222301017	3
5	2023-01-23 04:33:06.531831+07	2023-02-02 23:03:48.13527+07	13219033	Sidartha Prastya	f	4	Anggota TA222301017	3
1	2023-01-18 04:04:38.666893+07	2023-02-04 21:57:58.330849+07	22113344556677889900	Ipin Saipudin	t	1	adeknya phineas hehe	4
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

COPY public.rssi_logs (id, created_at, updated_at, rssi, lock_id, key_id, "timestamp", personel_id) FROM stdin;
1	2023-02-04 21:07:42.298464+07	2023-02-04 21:07:42.298464+07	204	26	1	2023-02-04 21:07:42.298254+07	1
2	2023-02-04 21:07:49.307961+07	2023-02-04 21:07:49.307961+07	204	26	1	2023-02-04 21:07:49.306932+07	1
3	2023-02-04 21:07:56.706552+07	2023-02-04 21:07:56.706552+07	204	26	1	2023-02-04 21:07:56.706029+07	1
4	2023-02-04 21:12:42.785829+07	2023-02-04 21:12:42.785829+07	204	26	3	2023-02-04 21:12:42.785643+07	4
5	2023-02-04 21:12:49.454143+07	2023-02-04 21:12:49.454143+07	204	26	3	2023-02-04 21:12:49.453631+07	4
6	2023-02-04 21:12:53.682003+07	2023-02-04 21:12:53.682003+07	204	26	3	2023-02-04 21:12:53.682003+07	4
7	2023-02-04 21:12:57.494272+07	2023-02-04 21:12:57.494272+07	204	26	3	2023-02-04 21:12:57.494272+07	4
8	2023-02-04 21:13:01.28402+07	2023-02-04 21:13:01.28402+07	204	26	3	2023-02-04 21:13:01.28402+07	4
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

SELECT pg_catalog.setval('public.access_logs_id_seq', 8, true);


--
-- Name: access_rules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.access_rules_id_seq', 54, true);


--
-- Name: healthcheck_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.healthcheck_logs_id_seq', 28, true);


--
-- Name: keys_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.keys_id_seq', 14, true);


--
-- Name: locks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.locks_id_seq', 26, true);


--
-- Name: personels_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.personels_id_seq', 6, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.roles_id_seq', 1, false);


--
-- Name: rssi_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.rssi_logs_id_seq', 8, true);


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

