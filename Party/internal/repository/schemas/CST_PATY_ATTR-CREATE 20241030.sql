create table cst_paty_attr (
	attr_row_id char(15) not null,
    paty_row_id char(15) not null,
	id_type char(2) not null,
	id_numb char(20) not null,
	bl_ext_id varchar(20)   ,
	ext_id varchar(70)   ,
	attr_name varchar(20) not null,
	attr_vlue varchar(40) ,
	crtd_dttm datetime year to second default year not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second default year not null,
	last_chng_by char(12)  not null,

primary key (attr_row_id)
    constraint cstp_paty_attr_01
);

--==============================================================
-- Index: cstp_paty_attr_01
--==============================================================
create unique index cstp_paty_attr_01 on cst_paty_attr (
	attr_row_id  ASC
);

--==============================================================
-- Index: cstf_paty_attr_02
--==============================================================
create index cstf_paty_attr_02 on cst_paty_attr (
	paty_row_id, attr_name  ASC
);

--==============================================================
-- Index: cstf_paty_attr_03
--==============================================================
create index cstf_paty_attr_03 on cst_paty_attr (
	id_numb, id_type, attr_name  ASC
);

--==============================================================
-- Index: cstf_paty_attr_04
--==============================================================
create index cstf_paty_attr_04 on cst_paty_attr (
	ext_id  ASC
);


--==============================================================
-- Index: cstf_paty_attr_05
--==============================================================
create index cstf_paty_attr_05 on cst_paty_attr (
	bl_ext_id  ASC
);