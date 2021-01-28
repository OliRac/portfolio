--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

-- Started on 2021-01-28 15:10:04

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
-- TOC entry 205 (class 1259 OID 16420)
-- Name: t_experience; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_experience (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    description text NOT NULL,
    company character varying(100) NOT NULL,
    lang_id integer NOT NULL,
    duration character varying(100) NOT NULL
);


ALTER TABLE public.t_experience OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 16418)
-- Name: t_experience_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_experience_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_experience_id_seq OWNER TO postgres;

--
-- TOC entry 3025 (class 0 OID 0)
-- Dependencies: 204
-- Name: t_experience_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_experience_id_seq OWNED BY public.t_experience.id;


--
-- TOC entry 2886 (class 2604 OID 16423)
-- Name: t_experience id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_experience ALTER COLUMN id SET DEFAULT nextval('public.t_experience_id_seq'::regclass);


--
-- TOC entry 3019 (class 0 OID 16420)
-- Dependencies: 205
-- Data for Name: t_experience; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_experience (id, title, description, company, lang_id, duration) FROM stdin;
1	Data Entry	Online part time work for a confidential project. Great opportunity to work with a full time college schedule.	Appen	1	May 2019 - December 2020
3	IT Technician	Assisted in the second revamp of the organization’s web site by remodeling the front-end. Coded multiple bash scripts to automate data fetching jobs. Produced documentation for business processes.	City of Westmount	1	May 2015 - November 2015
4	IT Technician	Aided in the revision of the organization’s entire web site; handled data migration and front-end development. Rewrote obsolete documentation for the department’s software and servers.	City of Westmount	1	May 2014 - November 2014
2	IT Technician	Built a project tracking tool on the organization’s intranet using HTML/CSS, JS, PHP and MariaDB. The tool allows anadministrator to add projects and data such as budgets, start dates and deadlines, and visualizes the data in multipletypes of graphs to aid in decision making. Set up and maintained a CentOS web server.	City of Westmount	1	May 2016 - August 2016
\.


--
-- TOC entry 3026 (class 0 OID 0)
-- Dependencies: 204
-- Name: t_experience_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.t_experience_id_seq', 4, true);


--
-- TOC entry 2887 (class 2606 OID 16427)
-- Name: t_experience lang_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_experience
    ADD CONSTRAINT lang_fk FOREIGN KEY (lang_id) REFERENCES public.t_language(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2021-01-28 15:10:04

--
-- PostgreSQL database dump complete
--

