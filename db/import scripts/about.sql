--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-28 15:07:44

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
-- TOC entry 215 (class 1259 OID 16498)
-- Name: t_about; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_about (
    id integer NOT NULL,
    quote character varying(100) NOT NULL,
    about text NOT NULL,
    title character varying(100) NOT NULL,
    lang_id integer NOT NULL
);


ALTER TABLE public.t_about OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16496)
-- Name: t_about_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_about_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_about_id_seq OWNER TO postgres;

--
-- TOC entry 3027 (class 0 OID 0)
-- Dependencies: 214
-- Name: t_about_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_about_id_seq OWNED BY public.t_about.id;


--
-- TOC entry 2886 (class 2604 OID 16501)
-- Name: t_about id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_about ALTER COLUMN id SET DEFAULT nextval('public.t_about_id_seq'::regclass);


--
-- TOC entry 3021 (class 0 OID 16498)
-- Dependencies: 215
-- Data for Name: t_about; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_about (id, quote, about, title, lang_id) FROM stdin;
1	Concentrate on the moment	Fresh Computer Science graduate. Looking to take my first step in the software development industry, whether in web, mobile, corporate or games. Started off in IT for the City of Westmount, but instead of doing support work, I was tasked with front-end development and creating small programs to make the department’s employee's lives easier. This provided motivation to pursue an education in Computer Science, deviating from IT / support.	Software Developer	1
\.


--
-- TOC entry 3028 (class 0 OID 0)
-- Dependencies: 214
-- Name: t_about_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_about_id_seq', 1, true);


--
-- TOC entry 2888 (class 2606 OID 16506)
-- Name: t_about t_about_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_about
    ADD CONSTRAINT t_about_pkey PRIMARY KEY (id);


--
-- TOC entry 2889 (class 2606 OID 16507)
-- Name: t_about lang_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_about
    ADD CONSTRAINT lang_fk FOREIGN KEY (lang_id) REFERENCES public.t_language(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2021-01-28 15:07:45

--
-- PostgreSQL database dump complete
--

