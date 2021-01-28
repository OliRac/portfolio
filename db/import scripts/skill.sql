--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-28 15:11:41

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
-- TOC entry 209 (class 1259 OID 16447)
-- Name: t_skill; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_skill (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.t_skill OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 16445)
-- Name: t_skill_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_skill_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_skill_id_seq OWNER TO postgres;

--
-- TOC entry 3026 (class 0 OID 0)
-- Dependencies: 208
-- Name: t_skill_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_skill_id_seq OWNED BY public.t_skill.id;


--
-- TOC entry 2886 (class 2604 OID 16450)
-- Name: t_skill id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_skill ALTER COLUMN id SET DEFAULT nextval('public.t_skill_id_seq'::regclass);


--
-- TOC entry 3020 (class 0 OID 16447)
-- Dependencies: 209
-- Data for Name: t_skill; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_skill (id, name) FROM stdin;
1	Python
2	C++
3	OpenGL
4	GLSL
5	Java
6	SQL
7	HTML/CSS
8	Javascript
9	Unity
10	Godot
\.


--
-- TOC entry 3027 (class 0 OID 0)
-- Dependencies: 208
-- Name: t_skill_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_skill_id_seq', 10, true);


--
-- TOC entry 2888 (class 2606 OID 16452)
-- Name: t_skill t_skill_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_skill
    ADD CONSTRAINT t_skill_pkey PRIMARY KEY (id);


-- Completed on 2021-01-28 15:11:41

--
-- PostgreSQL database dump complete
--

