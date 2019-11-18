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

DROP DATABASE IF EXISTS uic;
--
-- Name: uic; Type: DATABASE; Schema: -; Owner: falcon_plus
--

CREATE DATABASE uic WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE uic OWNER TO falcon_plus;

\connect uic

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
-- Name: rel_team_user; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.rel_team_user (
    id bigint NOT NULL,
    tid bigint NOT NULL,
    uid bigint NOT NULL
);


ALTER TABLE public.rel_team_user OWNER TO falcon_plus;

--
-- Name: rel_team_user_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.rel_team_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rel_team_user_id_seq OWNER TO falcon_plus;

--
-- Name: rel_team_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.rel_team_user_id_seq OWNED BY public.rel_team_user.id;


--
-- Name: session; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.session (
    id bigint NOT NULL,
    uid bigint NOT NULL,
    sig character varying(32) NOT NULL,
    expired bigint NOT NULL
);


ALTER TABLE public.session OWNER TO falcon_plus;

--
-- Name: session_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.session_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.session_id_seq OWNER TO falcon_plus;

--
-- Name: session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.session_id_seq OWNED BY public.session.id;


--
-- Name: team; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.team (
    id bigint NOT NULL,
    name character varying(64) NOT NULL,
    resume character varying(255) DEFAULT ''::character varying NOT NULL,
    creator bigint DEFAULT 0 NOT NULL,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.team OWNER TO falcon_plus;

--
-- Name: team_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.team_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.team_id_seq OWNER TO falcon_plus;

--
-- Name: team_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.team_id_seq OWNED BY public.team.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public."user" (
    id bigint NOT NULL,
    name character varying(64) NOT NULL,
    passwd character varying(64) DEFAULT ''::character varying NOT NULL,
    cnname character varying(128) DEFAULT ''::character varying NOT NULL,
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    phone character varying(16) DEFAULT ''::character varying NOT NULL,
    im character varying(32) DEFAULT ''::character varying NOT NULL,
    qq character varying(16) DEFAULT ''::character varying NOT NULL,
    role integer DEFAULT 0 NOT NULL,
    creator bigint DEFAULT 0 NOT NULL,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public."user" OWNER TO falcon_plus;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO falcon_plus;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.users_id_seq OWNED BY public."user".id;


--
-- Name: rel_team_user id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.rel_team_user ALTER COLUMN id SET DEFAULT nextval('public.rel_team_user_id_seq'::regclass);


--
-- Name: session id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.session ALTER COLUMN id SET DEFAULT nextval('public.session_id_seq'::regclass);


--
-- Name: team id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.team ALTER COLUMN id SET DEFAULT nextval('public.team_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: rel_team_user rel_team_user_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.rel_team_user
    ADD CONSTRAINT rel_team_user_pkey PRIMARY KEY (id);


--
-- Name: session session_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.session
    ADD CONSTRAINT session_pkey PRIMARY KEY (id);


--
-- Name: team team_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.team
    ADD CONSTRAINT team_pkey PRIMARY KEY (id);


--
-- Name: user users_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_rel_tid; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_rel_tid ON public.rel_team_user USING btree (tid);


--
-- Name: idx_rel_uid; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_rel_uid ON public.rel_team_user USING btree (uid);


--
-- Name: idx_session_sig; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_session_sig ON public.session USING btree (sig);


--
-- Name: idx_session_uid; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_session_uid ON public.session USING btree (uid);


--
-- Name: idx_team_name; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_team_name ON public.team USING btree (name);


--
-- Name: idx_user_name; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_user_name ON public."user" USING btree (name);


--
-- PostgreSQL database dump complete
--

