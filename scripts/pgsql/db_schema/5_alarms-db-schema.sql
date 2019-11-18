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

DROP DATABASE IF EXISTS alarms;
--
-- Name: alarms; Type: DATABASE; Schema: -; Owner: falcon_plus
--

CREATE DATABASE alarms WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE alarms OWNER TO falcon_plus;

\connect alarms

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
-- Name: event_cases; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.event_cases (
    id character varying(50),
    endpoint character varying(100) NOT NULL,
    metric character varying(200) NOT NULL,
    func character varying(50),
    cond character varying(200) NOT NULL,
    note character varying(500),
    max_step integer,
    current_step integer,
    priority integer NOT NULL,
    status character varying(20) NOT NULL,
    "timestamp" timestamp without time zone NOT NULL,
    update_at timestamp without time zone,
    closed_at timestamp without time zone,
    closed_note character varying(250),
    user_modified integer,
    tpl_creator character varying(64),
    expression_id integer,
    strategy_id integer,
    template_id integer,
    process_note integer,
    process_status character varying(20) DEFAULT 'unresolved'::character varying
);


ALTER TABLE public.event_cases OWNER TO falcon_plus;

--
-- Name: events; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.events (
    id bigint NOT NULL,
    event_caseid character varying(50),
    step integer,
    cond character varying(200) NOT NULL,
    status integer DEFAULT 0,
    "timestamp" timestamp without time zone
);


ALTER TABLE public.events OWNER TO falcon_plus;

--
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.events_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.events_id_seq OWNER TO falcon_plus;

--
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;


--
-- Name: events id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: idx_endpoint_strategy_id_emplate_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_endpoint_strategy_id_emplate_id ON public.event_cases USING btree (endpoint, strategy_id, template_id);


--
-- Name: idx_events_event_caseid; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_events_event_caseid ON public.events USING btree (event_caseid);


--
-- Name: pk_event_cases_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX pk_event_cases_id ON public.event_cases USING btree (id);


--
-- Name: events events_event_caseid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_event_caseid_fkey FOREIGN KEY (event_caseid) REFERENCES public.event_cases(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

