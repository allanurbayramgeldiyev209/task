--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Homebrew)
-- Dumped by pg_dump version 14.6 (Homebrew)

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

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: lessons; Type: TABLE; Schema: public; Owner: allanurbayramgeldiyev
--

CREATE TABLE public.lessons (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) DEFAULT NULL::character varying,
    user_id uuid NOT NULL,
    day_number smallint NOT NULL,
    order_number smallint NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.lessons OWNER TO allanurbayramgeldiyev;

--
-- Name: users; Type: TABLE; Schema: public; Owner: allanurbayramgeldiyev
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    phone_number character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO allanurbayramgeldiyev;

--
-- Data for Name: lessons; Type: TABLE DATA; Schema: public; Owner: allanurbayramgeldiyev
--

COPY public.lessons (id, name, user_id, day_number, order_number, created_at, updated_at, deleted_at) FROM stdin;
07aeeb74-ec04-4ef1-a936-8babd9a8dd6f	matematike	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	1	1	2023-08-10 12:56:43.006427+05	2023-08-10 12:56:43.006427+05	\N
58b47dea-b212-42f1-9a3d-b3a357c8a9e3	biologiya	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	1	2	2023-08-10 12:56:43.008905+05	2023-08-10 12:56:43.008905+05	\N
dc35f411-d931-47fe-8318-ba4281b7f768	bedenterbiye	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	1	3	2023-08-10 12:56:43.009501+05	2023-08-10 12:56:43.009501+05	\N
00bfb925-f026-4cb9-a837-e4fa4c10433a	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	1	4	2023-08-10 12:56:43.009968+05	2023-08-10 12:56:43.009968+05	\N
9cf08a52-3f03-457b-a433-0de8a6fc93fa	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	1	5	2023-08-10 12:56:43.010198+05	2023-08-10 12:56:43.010198+05	\N
b5babd89-5bae-4acc-9c17-3515e5906377	taryh	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	1	6	2023-08-10 12:56:43.010457+05	2023-08-10 12:56:43.010457+05	\N
a9d86320-16a8-4bed-9b71-f239626ecc08	taryh	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	2	1	2023-08-10 12:56:43.010689+05	2023-08-10 12:56:43.010689+05	\N
93784d9c-38fa-4065-ab90-6defec05e24c	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	2	2	2023-08-10 12:56:43.010873+05	2023-08-10 12:56:43.010873+05	\N
5d784771-a039-462c-8384-37eabbada4cf	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	2	3	2023-08-10 12:56:43.011064+05	2023-08-10 12:56:43.011064+05	\N
4bde9842-324e-4367-94b6-ebe6f6140ab9	yasayys durmus	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	2	4	2023-08-10 12:56:43.011242+05	2023-08-10 12:56:43.011242+05	\N
0530c916-690b-4a4c-8a00-95248f9a7a65	fizika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	2	5	2023-08-10 12:56:43.011417+05	2023-08-10 12:56:43.011417+05	\N
a3717a50-3684-4640-a06d-f42a5b84fe1a	synp sagat	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	2	6	2023-08-10 12:56:43.01159+05	2023-08-10 12:56:43.01159+05	\N
1d58c795-150c-49d4-9658-e1f529f20d79	fizika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	3	1	2023-08-10 12:56:43.011737+05	2023-08-10 12:56:43.011737+05	\N
2eb715df-2b85-4fa0-8016-865fa4ac8bcf	informatika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	3	2	2023-08-10 12:56:43.011909+05	2023-08-10 12:56:43.011909+05	\N
82f73be2-a3d6-44ee-8578-95cf65ff2896	bedenterbiye	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	3	3	2023-08-10 12:56:43.012056+05	2023-08-10 12:56:43.012056+05	\N
01fb896f-a2cf-4e5a-95ae-496891b23b96	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	3	4	2023-08-10 12:56:43.012213+05	2023-08-10 12:56:43.012213+05	\N
7177514c-53f2-4620-85a1-833fe54a1c90	informatika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	3	5	2023-08-10 12:56:43.012465+05	2023-08-10 12:56:43.012465+05	\N
071eeed5-70bd-454b-8cdd-cf8d278e7fce	inlis dili	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	3	6	2023-08-10 12:56:43.012719+05	2023-08-10 12:56:43.012719+05	\N
a6548940-1324-4fd0-98e1-dccbbdeeebb2	inlis dili	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	4	1	2023-08-10 12:56:43.012946+05	2023-08-10 12:56:43.012946+05	\N
e5f47095-abe2-4d93-a21f-7055a980e3f3	rus dili	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	4	2	2023-08-10 12:56:43.014516+05	2023-08-10 12:56:43.014516+05	\N
4f6b380f-1cc2-47c1-9aa0-843181d4830a	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	4	3	2023-08-10 12:56:43.014759+05	2023-08-10 12:56:43.014759+05	\N
126c105b-0719-438d-8204-a3ed298579a2	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	4	4	2023-08-10 12:56:43.015171+05	2023-08-10 12:56:43.015171+05	\N
7b29f32e-7358-43d8-a602-82a5a73446f2	matematika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	4	5	2023-08-10 12:56:43.015406+05	2023-08-10 12:56:43.015406+05	\N
f4f0883d-a97e-4d2b-bbbd-8482e6d364c8	turkmen dili	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	4	6	2023-08-10 12:56:43.015621+05	2023-08-10 12:56:43.015621+05	\N
7d2cdf11-18a7-462a-91bd-aba79f19be78	turkmen dili	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	5	1	2023-08-10 12:56:43.0158+05	2023-08-10 12:56:43.0158+05	\N
a12b633e-ab95-41c2-9962-062a001a644c	rus dili	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	5	2	2023-08-10 12:56:43.015981+05	2023-08-10 12:56:43.015981+05	\N
11a649dd-8371-4803-87aa-a34832f14f41	fizika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	5	3	2023-08-10 12:56:43.016189+05	2023-08-10 12:56:43.016189+05	\N
a7414b77-5c3e-4b5c-8929-4ec757c017fe	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	5	4	2023-08-10 12:56:43.016375+05	2023-08-10 12:56:43.016375+05	\N
d04bf9ce-3085-4c37-a4b9-d6631e7c8577	informatika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	5	5	2023-08-10 12:56:43.016553+05	2023-08-10 12:56:43.016553+05	\N
32913688-518f-4138-8f5e-bbc8baa4a452	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	5	6	2023-08-10 12:56:43.01673+05	2023-08-10 12:56:43.01673+05	\N
ccdac473-ffcf-416d-bf67-1a0dd2039f7b	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	6	1	2023-08-10 12:56:43.0169+05	2023-08-10 12:56:43.0169+05	\N
38ccb0a1-f6a8-4c2c-a278-a53b5e92662b	\N	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	6	2	2023-08-10 12:56:43.017092+05	2023-08-10 12:56:43.017092+05	\N
1a6dacb1-dda3-4432-a753-7aa7c72959bf	informatika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	6	3	2023-08-10 12:56:43.017278+05	2023-08-10 12:56:43.017278+05	\N
09de09fc-a2b4-4dea-990b-fd96eeb063b4	fizika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	6	4	2023-08-10 12:56:43.017501+05	2023-08-10 12:56:43.017501+05	\N
628655cb-bc5e-4660-948f-034ebea17fe3	biologiya	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	6	5	2023-08-10 12:56:43.017673+05	2023-08-10 12:56:43.017673+05	\N
374665c6-4520-4026-a000-0af1f272ba48	informatika	eb422a8e-3164-4709-8c70-7f23fdbd9cd3	6	6	2023-08-10 12:56:43.017866+05	2023-08-10 12:56:43.017866+05	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: allanurbayramgeldiyev
--

COPY public.users (id, phone_number, password, created_at, updated_at, deleted_at) FROM stdin;
eb422a8e-3164-4709-8c70-7f23fdbd9cd3	+99362420377	$2a$14$pEbO40AMmTo4II1wlG.ilu28Wd2s2W4kxQHLo1kLUIfwynjxtMZHS	2023-08-10 11:32:35.462055+05	2023-08-10 11:32:35.462055+05	\N
f93527b3-ecf5-4ed1-b52c-34afd83658c2	+99362658420	$2a$14$ytGIeTFIfWQBvKRdUmu9/.KJGrUSALVfkPMG29tzliJdObKR2CUrC	2023-08-10 11:56:11.242055+05	2023-08-10 11:56:11.242055+05	\N
\.


--
-- Name: lessons lessons_pkey; Type: CONSTRAINT; Schema: public; Owner: allanurbayramgeldiyev
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT lessons_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: allanurbayramgeldiyev
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: lessons users_lessons; Type: FK CONSTRAINT; Schema: public; Owner: allanurbayramgeldiyev
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT users_lessons FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

