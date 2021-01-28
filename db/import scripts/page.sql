--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-28 15:11:25

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
-- TOC entry 211 (class 1259 OID 16468)
-- Name: t_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_page (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.t_page OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16466)
-- Name: t_page_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_page_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_page_id_seq OWNER TO postgres;

--
-- TOC entry 3026 (class 0 OID 0)
-- Dependencies: 210
-- Name: t_page_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_page_id_seq OWNED BY public.t_page.id;


--
-- TOC entry 2886 (class 2604 OID 16471)
-- Name: t_page id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_page ALTER COLUMN id SET DEFAULT nextval('public.t_page_id_seq'::regclass);


--
-- TOC entry 3020 (class 0 OID 16468)
-- Dependencies: 211
-- Data for Name: t_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_page (id, name) FROM stdin;
1	resume
\.


--
-- TOC entry 3027 (class 0 OID 0)
-- Dependencies: 210
-- Name: t_page_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_page_id_seq', 1, true);


--
-- TOC entry 2888 (class 2606 OID 16473)
-- Name: t_page t_page_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_page
    ADD CONSTRAINT t_page_pkey PRIMARY KEY (id);


-- Completed on 2021-01-28 15:11:26

--
-- PostgreSQL database dump complete
--

