--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-28 15:10:21

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
-- TOC entry 213 (class 1259 OID 16476)
-- Name: t_header; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_header (
    id integer NOT NULL,
    section character varying(100) NOT NULL,
    lang_id integer NOT NULL,
    page_id integer NOT NULL
);


ALTER TABLE public.t_header OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16474)
-- Name: t_header_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_header_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_header_id_seq OWNER TO postgres;

--
-- TOC entry 3028 (class 0 OID 0)
-- Dependencies: 212
-- Name: t_header_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_header_id_seq OWNED BY public.t_header.id;


--
-- TOC entry 2886 (class 2604 OID 16479)
-- Name: t_header id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_header ALTER COLUMN id SET DEFAULT nextval('public.t_header_id_seq'::regclass);


--
-- TOC entry 3022 (class 0 OID 16476)
-- Dependencies: 213
-- Data for Name: t_header; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_header (id, section, lang_id, page_id) FROM stdin;
2	About	1	1
3	Skills	1	1
4	Knowledge	1	1
5	Education	1	1
6	Experience	1	1
7	Links	1	1
\.


--
-- TOC entry 3029 (class 0 OID 0)
-- Dependencies: 212
-- Name: t_header_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_header_id_seq', 7, true);


--
-- TOC entry 2888 (class 2606 OID 16481)
-- Name: t_header t_header_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_header
    ADD CONSTRAINT t_header_pkey PRIMARY KEY (id);


--
-- TOC entry 2889 (class 2606 OID 16482)
-- Name: t_header lang_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_header
    ADD CONSTRAINT lang_fk FOREIGN KEY (lang_id) REFERENCES public.t_language(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2890 (class 2606 OID 16487)
-- Name: t_header page_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_header
    ADD CONSTRAINT page_fk FOREIGN KEY (page_id) REFERENCES public.t_page(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2021-01-28 15:10:22

--
-- PostgreSQL database dump complete
--

