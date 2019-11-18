--
-- PostgreSQL database dump
--

-- Dumped from database version 11.0
-- Dumped by pg_dump version 11.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS dashboard;
--
-- Name: dashboard; Type: DATABASE; Schema: -; Owner: falcon_plus
--

CREATE DATABASE dashboard WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE dashboard OWNER TO falcon_plus;

\connect dashboard

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: dashboard_graph; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.dashboard_graph (
    id bigint NOT NULL,
    title character(128) NOT NULL,
    hosts character varying(10240) DEFAULT ''::character varying NOT NULL,
    counters character varying(1024) DEFAULT ''::character varying NOT NULL,
    screen_id bigint NOT NULL,
    timespan bigint DEFAULT 3600 NOT NULL,
    graph_type character(2) DEFAULT 'h'::bpchar NOT NULL,
    method character(8) DEFAULT ''::bpchar,
    "position" integer DEFAULT 0 NOT NULL,
    falcon_tags character varying(512) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.dashboard_graph OWNER TO falcon_plus;

--
-- Name: dashboard_graph_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.dashboard_graph_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dashboard_graph_id_seq OWNER TO falcon_plus;

--
-- Name: dashboard_graph_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.dashboard_graph_id_seq OWNED BY public.dashboard_graph.id;


--
-- Name: dashboard_screen; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.dashboard_screen (
    id bigint NOT NULL,
    pid bigint DEFAULT 0 NOT NULL,
    name character(128) NOT NULL,
    "time" timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.dashboard_screen OWNER TO falcon_plus;

--
-- Name: dashboard_screen_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.dashboard_screen_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dashboard_screen_id_seq OWNER TO falcon_plus;

--
-- Name: dashboard_screen_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.dashboard_screen_id_seq OWNED BY public.dashboard_screen.id;


--
-- Name: tmp_graph; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.tmp_graph (
    id bigint NOT NULL,
    endpoints character varying(10240) DEFAULT ''::character varying NOT NULL,
    counters character varying(10240) DEFAULT ''::character varying NOT NULL,
    ck character varying(32) NOT NULL,
    time_ timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.tmp_graph OWNER TO falcon_plus;

--
-- Name: tmp_graph_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.tmp_graph_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tmp_graph_id_seq OWNER TO falcon_plus;

--
-- Name: tmp_graph_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.tmp_graph_id_seq OWNED BY public.tmp_graph.id;


--
-- Name: dashboard_graph id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.dashboard_graph ALTER COLUMN id SET DEFAULT nextval('public.dashboard_graph_id_seq'::regclass);


--
-- Name: dashboard_screen id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.dashboard_screen ALTER COLUMN id SET DEFAULT nextval('public.dashboard_screen_id_seq'::regclass);


--
-- Name: tmp_graph id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.tmp_graph ALTER COLUMN id SET DEFAULT nextval('public.tmp_graph_id_seq'::regclass);


--
-- Name: dashboard_graph dashboard_graph_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.dashboard_graph
    ADD CONSTRAINT dashboard_graph_pkey PRIMARY KEY (id);


--
-- Name: dashboard_screen dashboard_screen_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.dashboard_screen
    ADD CONSTRAINT dashboard_screen_pkey PRIMARY KEY (id);


--
-- Name: tmp_graph tmp_graph_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.tmp_graph
    ADD CONSTRAINT tmp_graph_pkey PRIMARY KEY (id);


--
-- Name: idx_ck; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_ck ON public.tmp_graph USING btree (ck);


--
-- Name: idx_pid; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_pid ON public.dashboard_screen USING btree (pid);


--
-- Name: idx_pid_n; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_pid_n ON public.dashboard_screen USING btree (pid, name);


--
-- Name: idx_sid; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_sid ON public.dashboard_graph USING btree (screen_id);


--
-- PostgreSQL database dump complete
--

