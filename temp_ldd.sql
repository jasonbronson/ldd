--
-- PostgreSQL database dump
--

-- Dumped from database version 12.6 (Debian 12.6-1.pgdg100+1)
-- Dumped by pg_dump version 12.5 (Debian 12.5-1.pgdg100+1)

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
-- Name: logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.logs (
    id character varying NOT NULL,
    log_line character varying NOT NULL,
    last_error timestamp with time zone,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    matching_string character varying
);


ALTER TABLE public.logs OWNER TO postgres;

--
-- Name: logs_found; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.logs_found (
    logs_id character varying NOT NULL,
    time_start timestamp with time zone NOT NULL,
    time_end timestamp with time zone NOT NULL
);


ALTER TABLE public.logs_found OWNER TO admin;

--
-- Name: matches; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.matches (
    id character varying NOT NULL,
    matching_string character varying NOT NULL,
    name character varying,
    description character varying
);


ALTER TABLE public.matches OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Data for Name: logs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.logs (id, log_line, last_error, updated_at, matching_string) FROM stdin;
0ab817c3-2377-499e-a40f-20107432e5df	paypal_pay_job.go:50 @jobs.PaypalPayJob Cannot execute billing for billing agreement B-17E81760ES404101T 	2021-07-03 11:00:18+00	2021-07-04 00:38:36.169323+00	Cannot execute billing
\.


--
-- Data for Name: logs_found; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.logs_found (logs_id, time_start, time_end) FROM stdin;
0ab817c3-2377-499e-a40f-20107432e5df	2021-06-29 00:38:33+00	2021-07-04 00:38:33+00
\.


--
-- Data for Name: matches; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.matches (id, matching_string, name, description) FROM stdin;
78498515-e9c8-4c6f-958c-64b117d4e279	Cannot execute billing	paypal pay job	\N
1	cannot set transaction read-write	read only dbmode	\N
2	unable to unmarshal data	generic	\N
3	panic recovered	panic 	\N
4	error	error 	\N
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
20200615083569	t
\.


--
-- Name: logs logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (id);


--
-- Name: matches matches_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: logsfoundi; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX logsfoundi ON public.logs_found USING btree (logs_id);


--
-- Name: matchingstring; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX matchingstring ON public.logs USING btree (matching_string);


--
-- Name: logs_found logs_found_logs_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.logs_found
    ADD CONSTRAINT logs_found_logs_id_fkey FOREIGN KEY (logs_id) REFERENCES public.logs(id);


--
-- PostgreSQL database dump complete
--

