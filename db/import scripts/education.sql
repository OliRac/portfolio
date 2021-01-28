--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-28 15:09:48

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
-- TOC entry 203 (class 1259 OID 16407)
-- Name: t_education; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_education (
    id integer NOT NULL,
    degree character varying(100) NOT NULL,
    institution character varying(100) NOT NULL,
    location character varying(100) NOT NULL,
    grad_year date NOT NULL,
    lang_id integer NOT NULL
);


ALTER TABLE public.t_education OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 16405)
-- Name: t_education_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_education_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_education_id_seq OWNER TO postgres;

--
-- TOC entry 3027 (class 0 OID 0)
-- Dependencies: 202
-- Name: t_education_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_education_id_seq OWNED BY public.t_education.id;


--
-- TOC entry 2886 (class 2604 OID 16410)
-- Name: t_education id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_education ALTER COLUMN id SET DEFAULT nextval('public.t_education_id_seq'::regclass);


--
-- TOC entry 3021 (class 0 OID 16407)
-- Dependencies: 203
-- Data for Name: t_education; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_education (id, degree, institution, location, grad_year, lang_id) FROM stdin;
1	B.Sc Computer Science	Concordia University	Montréal	2020-12-01	1
2	DEC Informatique de gestion	CÉGEP du Vieux Montréal	Montréal	2016-12-01	1
3	BAC Science Informatique	Université Concordia	Montréal	2020-12-01	2
\.


--
-- TOC entry 3028 (class 0 OID 0)
-- Dependencies: 202
-- Name: t_education_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_education_id_seq', 3, true);


--
-- TOC entry 2888 (class 2606 OID 16412)
-- Name: t_education t_education_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_education
    ADD CONSTRAINT t_education_pkey PRIMARY KEY (id);


--
-- TOC entry 2889 (class 2606 OID 16413)
-- Name: t_education lang_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_education
    ADD CONSTRAINT lang_fk FOREIGN KEY (lang_id) REFERENCES public.t_language(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2021-01-28 15:09:48

--
-- PostgreSQL database dump complete
--

