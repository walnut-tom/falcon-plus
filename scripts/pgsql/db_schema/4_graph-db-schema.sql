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

DROP DATABASE IF EXISTS graph;
--
-- Name: graph; Type: DATABASE; Schema: -; Owner: falcon_plus
--

CREATE DATABASE graph WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE graph OWNER TO falcon_plus;

\connect graph

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
-- Name: endpoint; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.endpoint (
    id bigint NOT NULL,
    endpoint character varying(255) DEFAULT ''::character varying NOT NULL,
    ts integer,
    t_create date NOT NULL,
    t_modify timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.endpoint OWNER TO falcon_plus;

--
-- Name: endpoint_counter; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.endpoint_counter (
    id bigint NOT NULL,
    endpoint_id bigint NOT NULL,
    counter character varying(255) DEFAULT ''::character varying NOT NULL,
    step integer DEFAULT 60 NOT NULL,
    type character varying(16) NOT NULL,
    ts integer,
    t_create date NOT NULL,
    t_modify timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.endpoint_counter OWNER TO falcon_plus;

--
-- Name: endpoint_counter_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.endpoint_counter_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.endpoint_counter_id_seq OWNER TO falcon_plus;

--
-- Name: endpoint_counter_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.endpoint_counter_id_seq OWNED BY public.endpoint_counter.id;


--
-- Name: endpoint_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.endpoint_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.endpoint_id_seq OWNER TO falcon_plus;

--
-- Name: endpoint_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.endpoint_id_seq OWNED BY public.endpoint.id;


--
-- Name: tag_endpoint; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.tag_endpoint (
    id bigint NOT NULL,
    tag character varying(255) DEFAULT ''::character varying NOT NULL,
    endpoint_id bigint NOT NULL,
    ts bigint,
    t_create date NOT NULL,
    t_modify timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.tag_endpoint OWNER TO falcon_plus;

--
-- Name: tag_endpoint_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.tag_endpoint_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_endpoint_id_seq OWNER TO falcon_plus;

--
-- Name: tag_endpoint_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.tag_endpoint_id_seq OWNED BY public.tag_endpoint.id;


--
-- Name: endpoint id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.endpoint ALTER COLUMN id SET DEFAULT nextval('public.endpoint_id_seq'::regclass);


--
-- Name: endpoint_counter id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.endpoint_counter ALTER COLUMN id SET DEFAULT nextval('public.endpoint_counter_id_seq'::regclass);


--
-- Name: tag_endpoint id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.tag_endpoint ALTER COLUMN id SET DEFAULT nextval('public.tag_endpoint_id_seq'::regclass);


--
-- Name: endpoint_counter endpoint_counter_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.endpoint_counter
    ADD CONSTRAINT endpoint_counter_pkey PRIMARY KEY (id);


--
-- Name: endpoint endpoint_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.endpoint
    ADD CONSTRAINT endpoint_pkey PRIMARY KEY (id);


--
-- Name: tag_endpoint tag_endpoint_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.tag_endpoint
    ADD CONSTRAINT tag_endpoint_pkey PRIMARY KEY (id);


--
-- Name: idx_endpoint; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_endpoint ON public.endpoint USING btree (endpoint);


--
-- Name: idx_tag_endpoint_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_tag_endpoint_id ON public.tag_endpoint USING btree (tag, endpoint_id);


--
-- Name: indexidx_endpoint_id_counter; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX indexidx_endpoint_id_counter ON public.endpoint_counter USING btree (endpoint_id, counter);


--
-- PostgreSQL database dump complete
--

