create table CST_INAL_REGS (
	seqn serial  not null,
	cust_numb integer  not null,
	subr_numb char(12) not null,
	blpd_code char(2)  not null,
	mthd_code char(1)  not null,
	resn_code char(2)  ,
	auto_flag char(2) default 'N'  not null,
	actv_flag char(1)  not null,
	eai_dttm datetime year to second ,
	crtd_dttm datetime year to second not null,
	crtd_by char(12)  not null,
	last_chng_dttm datetime year to second not null,
	last_chng_by char(12)  not null,
  
primary key (seqn)
    constraint cstp_inal_regs_01
);

--==============================================================
-- Index: cstp_inal_regs_01
--==============================================================
create unique index cstp_inal_regs_01 on cst_inal_regs (
	seqn  ASC
);

--==============================================================
-- Index: cstn_inal_regs_02
--==============================================================
create index cstn_inal_regs_02 on cst_inal_regs (
	cust_numb, subr_numb ASC
);
  
--==============================================================
-- Index: cstn_inal_regs_03
--==============================================================
create index cstn_inal_regs_03 on cst_inal_regs (
	mthd_code ASC
);
