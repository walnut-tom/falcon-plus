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

DROP DATABASE IF EXISTS falcon_portal;
--
-- Name: falcon_portal; Type: DATABASE; Schema: -; Owner: falcon_plus
--

CREATE DATABASE falcon_portal WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE falcon_portal OWNER TO falcon_plus;

\connect falcon_portal

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
-- Name: action; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.action (
    id bigint NOT NULL,
    uic character varying(255) DEFAULT ''::character varying NOT NULL,
    url character varying(255) DEFAULT ''::character varying NOT NULL,
    callback integer DEFAULT 0 NOT NULL,
    before_callback_sms integer DEFAULT 0 NOT NULL,
    before_callback_mail integer DEFAULT 0 NOT NULL,
    after_callback_sms integer DEFAULT 0 NOT NULL,
    after_callback_mail integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.action OWNER TO falcon_plus;

--
-- Name: action_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.action_id_seq OWNER TO falcon_plus;

--
-- Name: action_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.action_id_seq OWNED BY public.action.id;


--
-- Name: alert_link; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.alert_link (
    id bigint NOT NULL,
    path character varying(16) DEFAULT ''::character varying NOT NULL,
    content text NOT NULL,
    create_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.alert_link OWNER TO falcon_plus;

--
-- Name: alert_link_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.alert_link_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.alert_link_id_seq OWNER TO falcon_plus;

--
-- Name: alert_link_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.alert_link_id_seq OWNED BY public.alert_link.id;


--
-- Name: cluster; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.cluster (
    id bigint NOT NULL,
    grp_id bigint NOT NULL,
    numerator character varying(10240) NOT NULL,
    denominator character varying(10240) NOT NULL,
    endpoint character varying(255) NOT NULL,
    metric character varying(255) NOT NULL,
    tags character varying(255) NOT NULL,
    ds_type character varying(255) NOT NULL,
    step integer NOT NULL,
    last_update timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    creator character varying(255) NOT NULL
);


ALTER TABLE public.cluster OWNER TO falcon_plus;

--
-- Name: cluster_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.cluster_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cluster_id_seq OWNER TO falcon_plus;

--
-- Name: cluster_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.cluster_id_seq OWNED BY public.cluster.id;


--
-- Name: expression; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.expression (
    id bigint DEFAULT nextval('expression_id_seq'::regclass) PRIMARY KEY,
    expression character varying(1024) NOT NULL,
    func character varying(16) DEFAULT 'all(#1)'::character varying NOT NULL,
    op character varying(8) DEFAULT ''::character varying NOT NULL,
    right_value character varying(16) DEFAULT ''::character varying NOT NULL,
    max_step integer DEFAULT 1 NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    note character varying(1024) DEFAULT ''::character varying NOT NULL,
    action_id bigint DEFAULT '0'::bigint NOT NULL,
    create_user character varying(64) DEFAULT ''::character varying NOT NULL,
    pause integer DEFAULT 0 NOT NULL
);


-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX expression_pkey ON expression(id int8_ops);



ALTER TABLE public.expression OWNER TO falcon_plus;

--
-- Name: grp; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.grp (
    id bigint NOT NULL,
    grp_name character varying(255) DEFAULT ''::character varying NOT NULL,
    create_user character varying(64) DEFAULT ''::character varying NOT NULL,
    create_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    come_from integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.grp OWNER TO falcon_plus;

--
-- Name: grp_host; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.grp_host (
    grp_id bigint NOT NULL,
    host_id bigint NOT NULL
);


ALTER TABLE public.grp_host OWNER TO falcon_plus;

--
-- Name: grp_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.grp_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.grp_id_seq OWNER TO falcon_plus;

--
-- Name: grp_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.grp_id_seq OWNED BY public.grp.id;


--
-- Name: grp_tpl; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.grp_tpl (
    grp_id bigint NOT NULL,
    tpl_id bigint NOT NULL,
    bind_user character varying(64) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.grp_tpl OWNER TO falcon_plus;

--
-- Name: host; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.host (
    id bigint NOT NULL,
    hostname character varying(255) DEFAULT ''::character varying NOT NULL,
    ip character varying(16) DEFAULT ''::character varying NOT NULL,
    agent_version character varying(16) DEFAULT ''::character varying NOT NULL,
    plugin_version character varying(128) DEFAULT ''::character varying NOT NULL,
    maintain_begin bigint DEFAULT 0 NOT NULL,
    maintain_end bigint DEFAULT 0 NOT NULL,
    update_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.host OWNER TO falcon_plus;

--
-- Name: host_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.host_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.host_id_seq OWNER TO falcon_plus;

--
-- Name: host_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.host_id_seq OWNED BY public.host.id;


--
-- Name: mockcfg; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.mockcfg (
    id bigint NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    obj character varying(10240) DEFAULT ''::character varying NOT NULL,
    obj_type character varying(255) DEFAULT ''::character varying NOT NULL,
    metric character varying(128) DEFAULT ''::character varying NOT NULL,
    tags character varying(1024) DEFAULT ''::character varying NOT NULL,
    dstype character varying(32) DEFAULT 'GAUGE'::character varying NOT NULL,
    step integer DEFAULT 60 NOT NULL,
    mock numeric DEFAULT 0 NOT NULL,
    creator character varying(64) DEFAULT ''::character varying NOT NULL,
    t_create date DEFAULT CURRENT_DATE NOT NULL,
    t_modify timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.mockcfg OWNER TO falcon_plus;

--
-- Name: mockcfg_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.mockcfg_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mockcfg_id_seq OWNER TO falcon_plus;

--
-- Name: mockcfg_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.mockcfg_id_seq OWNED BY public.mockcfg.id;


--
-- Name: plugin_dir; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.plugin_dir (
    id bigint NOT NULL,
    grp_id bigint NOT NULL,
    dir character varying(255) NOT NULL,
    create_user character varying(64) DEFAULT ''::character varying NOT NULL,
    create_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.plugin_dir OWNER TO falcon_plus;

--
-- Name: plugin_dir_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.plugin_dir_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.plugin_dir_id_seq OWNER TO falcon_plus;

--
-- Name: plugin_dir_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.plugin_dir_id_seq OWNED BY public.plugin_dir.id;


--
-- Name: strategy; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.strategy (
    id bigint NOT NULL,
    metric character varying(128) DEFAULT ''::character varying NOT NULL,
    tags character varying(256) DEFAULT ''::character varying NOT NULL,
    max_step integer DEFAULT 1 NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    func character varying(16) DEFAULT 'all(#1)'::character varying NOT NULL,
    op character varying(8) DEFAULT ''::character varying NOT NULL,
    right_value character varying(64) NOT NULL,
    note character varying(128) DEFAULT ''::character varying NOT NULL,
    run_begin character varying(16) DEFAULT ''::character varying NOT NULL,
    run_end character varying(16) DEFAULT ''::character varying NOT NULL,
    tpl_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE public.strategy OWNER TO falcon_plus;

--
-- Name: strategy_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.strategy_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.strategy_id_seq OWNER TO falcon_plus;

--
-- Name: strategy_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.strategy_id_seq OWNED BY public.strategy.id;


--
-- Name: tpl; Type: TABLE; Schema: public; Owner: falcon_plus
--

CREATE TABLE public.tpl (
    id bigint NOT NULL,
    tpl_name character varying(255) DEFAULT ''::character varying NOT NULL,
    parent_id bigint DEFAULT 0 NOT NULL,
    action_id bigint DEFAULT 0 NOT NULL,
    create_user character varying(64) DEFAULT ''::character varying NOT NULL,
    create_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.tpl OWNER TO falcon_plus;

--
-- Name: tpl_id_seq; Type: SEQUENCE; Schema: public; Owner: falcon_plus
--

CREATE SEQUENCE public.tpl_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tpl_id_seq OWNER TO falcon_plus;

--
-- Name: tpl_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: falcon_plus
--

ALTER SEQUENCE public.tpl_id_seq OWNED BY public.tpl.id;


--
-- Name: action id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.action ALTER COLUMN id SET DEFAULT nextval('public.action_id_seq'::regclass);


--
-- Name: alert_link id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.alert_link ALTER COLUMN id SET DEFAULT nextval('public.alert_link_id_seq'::regclass);


--
-- Name: cluster id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.cluster ALTER COLUMN id SET DEFAULT nextval('public.cluster_id_seq'::regclass);


--
-- Name: grp id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.grp ALTER COLUMN id SET DEFAULT nextval('public.grp_id_seq'::regclass);


--
-- Name: host id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.host ALTER COLUMN id SET DEFAULT nextval('public.host_id_seq'::regclass);


--
-- Name: mockcfg id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.mockcfg ALTER COLUMN id SET DEFAULT nextval('public.mockcfg_id_seq'::regclass);


--
-- Name: plugin_dir id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.plugin_dir ALTER COLUMN id SET DEFAULT nextval('public.plugin_dir_id_seq'::regclass);


--
-- Name: strategy id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.strategy ALTER COLUMN id SET DEFAULT nextval('public.strategy_id_seq'::regclass);


--
-- Name: tpl id; Type: DEFAULT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.tpl ALTER COLUMN id SET DEFAULT nextval('public.tpl_id_seq'::regclass);


--
-- Name: action action_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.action
    ADD CONSTRAINT action_pkey PRIMARY KEY (id);


--
-- Name: alert_link alert_link_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.alert_link
    ADD CONSTRAINT alert_link_pkey PRIMARY KEY (id);


--
-- Name: cluster cluster_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.cluster
    ADD CONSTRAINT cluster_pkey PRIMARY KEY (id);


--
-- Name: grp_host grp_host_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.grp_host
    ADD CONSTRAINT grp_host_pkey PRIMARY KEY (grp_id, host_id);


--
-- Name: grp grp_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.grp
    ADD CONSTRAINT grp_pkey PRIMARY KEY (id);


--
-- Name: grp_tpl grp_tpl_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.grp_tpl
    ADD CONSTRAINT grp_tpl_pkey PRIMARY KEY (grp_id, tpl_id);


--
-- Name: host host_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.host
    ADD CONSTRAINT host_pkey PRIMARY KEY (id);


--
-- Name: mockcfg mockcfg_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.mockcfg
    ADD CONSTRAINT mockcfg_pkey PRIMARY KEY (id);


--
-- Name: plugin_dir plugin_dir_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.plugin_dir
    ADD CONSTRAINT plugin_dir_pkey PRIMARY KEY (id);


--
-- Name: strategy strategy_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.strategy
    ADD CONSTRAINT strategy_pkey PRIMARY KEY (id);


--
-- Name: tpl tpl_pkey; Type: CONSTRAINT; Schema: public; Owner: falcon_plus
--

ALTER TABLE ONLY public.tpl
    ADD CONSTRAINT tpl_pkey PRIMARY KEY (id);


--
-- Name: alert_path; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX alert_path ON public.alert_link USING btree (path);


--
-- Name: idx_grp_host_grp_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_grp_host_grp_id ON public.grp_host USING btree (grp_id);


--
-- Name: idx_grp_host_host_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_grp_host_host_id ON public.grp_host USING btree (host_id);


--
-- Name: idx_grp_tpl_grp_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_grp_tpl_grp_id ON public.grp_tpl USING btree (grp_id);


--
-- Name: idx_grp_tpl_tpl_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_grp_tpl_tpl_id ON public.grp_tpl USING btree (tpl_id);


--
-- Name: idx_host_grp_grp_name; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_host_grp_grp_name ON public.grp USING btree (grp_name);


--
-- Name: idx_host_hostname; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_host_hostname ON public.host USING btree (hostname);


--
-- Name: idx_plugin_dir_grp_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_plugin_dir_grp_id ON public.plugin_dir USING btree (grp_id);


--
-- Name: idx_strategy_tpl_id; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_strategy_tpl_id ON public.strategy USING btree (tpl_id);


--
-- Name: idx_tpl_create_user; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE INDEX idx_tpl_create_user ON public.tpl USING btree (create_user);


--
-- Name: idx_tpl_name; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX idx_tpl_name ON public.tpl USING btree (tpl_name);


--
-- Name: uniq_name; Type: INDEX; Schema: public; Owner: falcon_plus
--

CREATE UNIQUE INDEX uniq_name ON public.mockcfg USING btree (name);


--
-- PostgreSQL database dump complete
--

