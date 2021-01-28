PGDMP     7    
                 y            portfolio_dev    13.1    13.1     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16396    portfolio_dev    DATABASE     q   CREATE DATABASE portfolio_dev WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE portfolio_dev;
                postgres    false            �           0    0    DATABASE portfolio_dev    COMMENT     h   COMMENT ON DATABASE portfolio_dev IS 'Will mostly contain strings of text to be used on my portfolio.';
                   postgres    false    3027            �            1259    16434    t_knowledge    TABLE     �   CREATE TABLE public.t_knowledge (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    lang_id integer NOT NULL
);
    DROP TABLE public.t_knowledge;
       public         heap    postgres    false            �            1259    16432    t_knowledge_id_seq    SEQUENCE     �   CREATE SEQUENCE public.t_knowledge_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.t_knowledge_id_seq;
       public          postgres    false    207            �           0    0    t_knowledge_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.t_knowledge_id_seq OWNED BY public.t_knowledge.id;
          public          postgres    false    206            F           2604    16437    t_knowledge id    DEFAULT     p   ALTER TABLE ONLY public.t_knowledge ALTER COLUMN id SET DEFAULT nextval('public.t_knowledge_id_seq'::regclass);
 =   ALTER TABLE public.t_knowledge ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    207    206    207            �          0    16434    t_knowledge 
   TABLE DATA           8   COPY public.t_knowledge (id, name, lang_id) FROM stdin;
    public          postgres    false    207   5       �           0    0    t_knowledge_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.t_knowledge_id_seq', 6, true);
          public          postgres    false    206            H           2606    16439    t_knowledge t_knowledge_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.t_knowledge
    ADD CONSTRAINT t_knowledge_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.t_knowledge DROP CONSTRAINT t_knowledge_pkey;
       public            postgres    false    207            I           2606    16440    t_knowledge lang_fk    FK CONSTRAINT     �   ALTER TABLE ONLY public.t_knowledge
    ADD CONSTRAINT lang_fk FOREIGN KEY (lang_id) REFERENCES public.t_language(id) ON UPDATE CASCADE ON DELETE CASCADE;
 =   ALTER TABLE ONLY public.t_knowledge DROP CONSTRAINT lang_fk;
       public          postgres    false    207            �   �   x�̻�0���y�����mG��*� L,&2�"u�$��ǌ�?���9��8Q`!h�
NT�1�YF;�vnyi�\�5tX񁅊��s�'{�`{��$�Ol�����(�qm�ҧR�7.E�}a���z(     